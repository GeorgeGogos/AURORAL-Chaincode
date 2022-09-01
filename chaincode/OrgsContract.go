package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/state"

	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"

	"fmt"
	"time"

	"github.com/s7techlab/cckit/router"
)

func CreateContract(c router.Context) (interface{}, error) {
	contractPayload := c.Param("contractPayload").(payload.ContractPayload) // Assert the chaincode parameter

	logging.CCLoggerInstance.Printf("Received input: %s. Attempting to validate contract request...\n", contractPayload.String())
	if err := contractPayload.Validate(); err != nil {
		retErr := fmt.Errorf("Error: Validate() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}

	for i := 0; i < len(contractPayload.Items); i++ {
		if err := contractPayload.Items[i].Validate(); err != nil {
			retErr := fmt.Errorf("Error: Validate() returned error: %s", err.Error())
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
	}

	contractState := state.ContractState{
		ContractId:     contractPayload.ContractId,
		ContractType:   contractPayload.ContractType,
		ContractStatus: contractPayload.ContractStatus,
		Orgs:           contractPayload.Orgs,
		Items:          contractPayload.Items,
		LastUpdated:    time.Now(),
		Created:        time.Now(),
	}

	logging.CCLoggerInstance.Printf("CreateContract function invokes chaincode. Quoted Output: %s\n", contractState.String())

	if err := c.State().Insert(contractState); err != nil {
		retErr := fmt.Errorf("Error while attempting to insert Contract info to state: %s", err)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, err
	}

	return nil, nil
}
