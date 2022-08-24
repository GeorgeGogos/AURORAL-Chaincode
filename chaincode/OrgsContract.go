package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/state"

	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"

	"fmt"

	"github.com/s7techlab/cckit/router"
)

func CreateContract(c router.Context) (interface{}, error) {
	var (
		contractPayload = c.Param("contractPayload").(payload.ContractPayload) // Assert the chaincode parameter
		contractState   = &state.ContractState{
			ContractId:     contractPayload.ContractId,
			ContractType:   contractPayload.ContractType,
			ContractStatus: contractPayload.ContractStatus,
			Orgs:           contractPayload.Orgs,
			Items:          contractPayload.Items,
			LastUpdated:    contractPayload.LastUpdated,
			Created:        contractPayload.Created,
		}
	)
	logging.CCLoggerInstance.Printf("CreateContract function invokes chaincode. Output: %v\n", contractState)
	if err := c.State().Insert(contractState); err != nil {
		retErr := fmt.Errorf("error while attempting to insert Contract info to state: %s", err)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, err
	}
	logging.CCLoggerInstance.Printf("Successfully initialized Account chaincode!")
	return nil, nil
}
