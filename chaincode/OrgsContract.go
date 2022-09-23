package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/state"

	"fmt"

	logging "github.com/CERTH-ITI-DLT-Lab/hlf-cc-logging"

	"github.com/s7techlab/cckit/router"
)

func ProposeContract(c router.Context) (interface{}, error) {
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
	if owner, err := onlyContractOrgs(c); err != nil {
		retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
		return nil, retErr
	} else if owner != string(contractPayload.Orgs[0]) && owner != string(contractPayload.Orgs[1]) {
		retErr := fmt.Errorf("The Org invoking the chaincode does not match the Orgs in payload")
		return nil, retErr
	}

	stateStub := state.NewStateStub(c)
	if err := stateStub.NewContract(contractPayload); err != nil {
		retErr := fmt.Errorf("Error: CreateContract returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}
	logging.CCLoggerInstance.Printf("Successfully created a Contract between Orgs: %s, %s", contractPayload.Orgs[0], contractPayload.Orgs[1])

	return nil, nil
}

func AcceptContract(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	logging.CCLoggerInstance.Printf("Received input: %s. Attempting to validate contract request...\n", contractID)

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}); err != nil {

		retErr := fmt.Errorf("The requested Contract does not exists: %s", err.Error())
		fmt.Printf("Type of stateContract: %#v\n", stateContract)
		return nil, retErr
	} /*else if stateContract.ContractStatus == "Rejected" || stateContract.ContractStatus == "Accepted" {
		retErr := fmt.Errorf("Error in Contract payload: %s", err.Error())
		return nil, retErr
	} else {
		stateContract.ContractStatus = "Accepted"
	}*/

	return nil, nil
}
