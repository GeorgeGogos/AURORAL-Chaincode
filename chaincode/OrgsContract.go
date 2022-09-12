package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/state"

	"fmt"

	logging "github.com/CERTH-ITI-DLT-Lab/hlf-cc-logging"

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
		if err := contractPayload.Items[i].Validate(contractPayload); err != nil {
			retErr := fmt.Errorf("Error: Validate() returned error: %s", err.Error())
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
	}

	logging.CCLoggerInstance.Printf("Checking ACL rules\n")
	if owner := onlyContractOrgs(c); owner != string(contractPayload.Orgs[0]) && owner != string(contractPayload.Orgs[1]) {
		retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL")
		return nil, retErr
	}

	stateStub := state.NewStateStub(c)
	if err := stateStub.NewContract(contractPayload); err != nil {
		retErr := fmt.Errorf("Error: CreateContract returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}

	return nil, nil
}
