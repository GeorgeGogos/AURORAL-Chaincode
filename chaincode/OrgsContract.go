package chaincode

import (
	"encoding/json"

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
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
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

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {

		retErr := fmt.Errorf("The requested Contract does not exists: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		fmt.Printf("Data of stateContract: %s\n", stateContract)
		acceptedContract := stateContract.(state.ContractState)

		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else if owner != string(acceptedContract.Orgs[0]) && owner != string(acceptedContract.Orgs[1]) {
			retErr := fmt.Errorf("The Org invoking the chaincode does not match the Orgs in payload")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
		if acceptedContract.ContractStatus == "Rejected" || acceptedContract.ContractStatus == "Accepted" {
			retErr := fmt.Errorf("Error in Contract payload.Contract Status is not 'Pending'.")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			acceptedContract.ContractStatus = "Accepted"
			fmt.Printf("Data of acceptedContract: %s\n", acceptedContract)
			if err := c.State().Put(acceptedContract); err != nil {
				retErr := fmt.Errorf("Error: Put() returned error: %s", err.Error())
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			}

		}

	}

	return nil, nil
}

func RejectContract(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	logging.CCLoggerInstance.Printf("Received input: %s. Attempting to validate contract request...\n", contractID)

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {

		retErr := fmt.Errorf("The requested Contract does not exists: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		fmt.Printf("Data of stateContract: %s\n", stateContract)
		rejectedContract := stateContract.(state.ContractState)

		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else if owner != string(rejectedContract.Orgs[0]) && owner != string(rejectedContract.Orgs[1]) {
			retErr := fmt.Errorf("The Org invoking the chaincode does not match the Orgs in payload")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}

		if rejectedContract.ContractStatus == "Rejected" || rejectedContract.ContractStatus == "Accepted" {
			retErr := fmt.Errorf("Error in Contract payload. Contract Status is not 'Pending'.")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			rejectedContract.ContractStatus = "Rejected"
			fmt.Printf("Data of rejectedContract: %s\n", rejectedContract)
			if err := c.State().Put(rejectedContract); err != nil {
				retErr := fmt.Errorf("Error: Put() returned error: %s", err.Error())
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			}

		}

	}

	return nil, nil
}

func DissolveContract(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")

	if exists, err := c.State().Exists(state.ContractState{ContractId: contractID}); err != nil {
		retErr := fmt.Errorf("Error: Exists() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else if !exists {
		retErr := fmt.Errorf("Error: Invalid delete operation, contract with ID: %s does not exist in contract state", contractID)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		stateContract, _ := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{})
		deletedContract := stateContract.(state.ContractState)

		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else if owner != string(deletedContract.Orgs[0]) && owner != string(deletedContract.Orgs[1]) {
			retErr := fmt.Errorf("The Org invoking the chaincode does not match the Orgs in payload")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
	}

	if err := c.State().Delete(&state.ContractState{ContractId: contractID}); err != nil {
		retErr := fmt.Errorf("Error: Delete() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr

	}

	return nil, nil
}

func GetContractByID(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	if owner, err := onlyContractOrgs(c); err != nil {
		retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
		return nil, retErr
	} else {
		if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {
			retErr := fmt.Errorf("The requested Contract does not exist: %s", err.Error())
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			stateContractStruct := stateContract.(state.ContractState)
			if owner != string(stateContractStruct.Orgs[0]) && owner != string(stateContractStruct.Orgs[1]) {
				retErr := fmt.Errorf("The Org invoking the chaincode does not match the Orgs in payload")
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			} else {
				if stateContractStruct.ContractStatus != "Accepted" {
					retErr := fmt.Errorf("There is no contract with this ID")
					logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
					return nil, retErr
				} else {
					logging.CCLoggerInstance.Printf("Attemting to marshal output...\n")
					marshaledOutput, err := json.Marshal(stateContractStruct)
					if err != nil {
						retErr := fmt.Errorf("error: json.Marshal() of output keys returned error: %s", err.Error())
						logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
						return nil, retErr
					}
					logging.CCLoggerInstance.Printf("Query successfully completed! Returning output: %s\n", string(marshaledOutput))
					return marshaledOutput, nil
				}

			}

		}
	}
}

func GetContracts(c router.Context) (interface{}, error) {
	if querylist, err := c.State().List([]string{state.ContractStateEntity}, &state.ContractState{}); err != nil {
		retErr := fmt.Errorf("Error: List() returned error in list function: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else {

			queriedInterfaceArray := querylist.([]interface{})
			var outputList []state.ContractState
			for _, curQueriedObj := range queriedInterfaceArray {
				stateContractStruct := curQueriedObj.(state.ContractState)
				if (owner == string(stateContractStruct.Orgs[0]) || owner == string(stateContractStruct.Orgs[1])) && stateContractStruct.ContractStatus == "Accepted" {
					outputList = append(outputList, stateContractStruct)
				}
			}
			if len(outputList) == 0 {
				emptyResultArray := make([]state.ContractState, 0)
				logging.CCLoggerInstance.Printf("Query successfully completed! Returning output: %v\n", emptyResultArray)
				return emptyResultArray, nil
			}
			logging.CCLoggerInstance.Printf("Attemting to marshal output...\n")
			marshaledOutput, err := json.Marshal(outputList)
			if err != nil {
				retErr := fmt.Errorf("error: json.Marshal() of output keys returned error: %s", err.Error())
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			}
			logging.CCLoggerInstance.Printf("Query successfully completed! Returning output: %s\n", string(marshaledOutput))
			return marshaledOutput, nil
		}

	}
}

func GetContractIDs(c router.Context) (interface{}, error) {
	if querylist, err := c.State().List([]string{state.ContractStateEntity}, &state.ContractState{}); err != nil {
		retErr := fmt.Errorf("Error: List() returned error in list function: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else {

			queriedInterfaceArray := querylist.([]interface{})
			var outputList []string
			for _, curQueriedObj := range queriedInterfaceArray {
				stateContractStruct := curQueriedObj.(state.ContractState)
				if (owner == string(stateContractStruct.Orgs[0]) || owner == string(stateContractStruct.Orgs[1])) && stateContractStruct.ContractStatus == "Accepted" {
					outputList = append(outputList, stateContractStruct.ContractId)
				}
			}
			if len(outputList) == 0 {
				emptyResultArray := make([]string, 0)
				logging.CCLoggerInstance.Printf("Query successfully completed! Returning output: %v\n", emptyResultArray)
				return emptyResultArray, nil
			}
			logging.CCLoggerInstance.Printf("Attemting to marshal output...\n")
			marshaledOutput, err := json.Marshal(outputList)
			if err != nil {
				retErr := fmt.Errorf("error: json.Marshal() of output keys returned error: %s", err.Error())
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			}
			logging.CCLoggerInstance.Printf("Query successfully completed! Returning output: %s\n", string(marshaledOutput))
			return marshaledOutput, nil
		}

	}
}

func UpdateContractItem(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	itemPayload := c.Param("itemPayload").(payload.Item) // Assert the chaincode parameter

	logging.CCLoggerInstance.Printf("Checking Contract Exists()\n")
	if exists, err := c.State().Exists(state.ContractState{ContractId: contractID}); err != nil {
		retErr := fmt.Errorf("Error: Exists() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else if !exists {
		retErr := fmt.Errorf("Error: Invalid delete operation, contract with ID: %s does not exist in contract state", contractID)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {
		retErr := fmt.Errorf("The requested Contract does not exist: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		stateContractStruct := stateContract.(state.ContractState)
		stateToPayloadStruct := payload.ContractPayload{
			ContractId:   stateContractStruct.ContractId,
			ContractType: stateContractStruct.ContractType,
			Orgs:         stateContractStruct.Orgs,
			Items:        stateContractStruct.Items,
		}
		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else if owner != string(itemPayload.OrgId) || stateContractStruct.ContractStatus != "Accepted" {
			retErr := fmt.Errorf("The Org invoking the chaincode does not match the Item's Org")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
		logging.CCLoggerInstance.Printf("Received input: %s. Attempting to validate contract request...\n", itemPayload.String())
		if err := itemPayload.Validate(stateToPayloadStruct); err != nil {
			retErr := fmt.Errorf("Error: Validate() returned error: %s", err.Error())
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
		logging.CCLoggerInstance.Printf("Checking Items array length\n")
		if len(stateContractStruct.Items) == 0 {
			retErr := fmt.Errorf("The requested Contract does not include any Items")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			itemFound := false
			for i, curQueriedObj := range stateContractStruct.Items {
				if curQueriedObj.ObjectId == itemPayload.ObjectId && curQueriedObj.UnitId == itemPayload.UnitId && curQueriedObj.OrgId == itemPayload.OrgId && curQueriedObj.ObjectType == itemPayload.ObjectType {
					itemFound = true
					curQueriedObj.Enabled = itemPayload.Enabled
					curQueriedObj.Write = itemPayload.Write
					stateContractStruct.Items[i] = curQueriedObj
				}
			}
			if !itemFound {
				retErr := fmt.Errorf("The requested Contract does not include the requested Item")
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			} else {
				if err := c.State().Put(stateContractStruct); err != nil {
					retErr := fmt.Errorf("Error: Put() returned error: %s", err.Error())
					return nil, retErr
				}
				fmt.Printf("Updated Contract: %s\n", &stateContractStruct)
				return stateContractStruct, nil
			}
		}
	}
}

func DeleteContractItem(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	itemID := c.ParamString("item_ID")

	logging.CCLoggerInstance.Printf("Checking Contract Exists()\n")
	if exists, err := c.State().Exists(state.ContractState{ContractId: contractID}); err != nil {
		retErr := fmt.Errorf("Error: Exists() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else if !exists {
		retErr := fmt.Errorf("Error: Invalid delete operation, contract with ID: %s does not exist in contract state", contractID)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {
		retErr := fmt.Errorf("The requested Contract does not exist: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		stateContractStruct := stateContract.(state.ContractState)

		logging.CCLoggerInstance.Printf("Checking Items array length\n")
		if len(stateContractStruct.Items) == 0 {
			retErr := fmt.Errorf("The requested Contract does not include any Items")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			itemFound := false
			for i, curQueriedObj := range stateContractStruct.Items {
				if curQueriedObj.ObjectId == itemID {
					itemFound = true
					logging.CCLoggerInstance.Printf("Checking ACL rules\n")
					if owner, err := onlyContractOrgs(c); err != nil {
						retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
						return nil, retErr
					} else if owner != string(curQueriedObj.OrgId) || stateContractStruct.ContractStatus != "Accepted" {
						retErr := fmt.Errorf("The Org invoking the chaincode does not match the Item's Org")
						logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
						return nil, retErr
					}
					stateContractStruct.Items[i] = stateContractStruct.Items[len(stateContractStruct.Items)-1]
					stateContractStruct.Items[len(stateContractStruct.Items)-1] = payload.Item{}
					stateContractStruct.Items = stateContractStruct.Items[0 : len(stateContractStruct.Items)-1]
					break
				}
			}
			if !itemFound {
				retErr := fmt.Errorf("The requested Contract does not include the requested Item")
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			} else {
				if err := c.State().Put(stateContractStruct); err != nil {
					retErr := fmt.Errorf("Error: Put() returned error: %s", err.Error())
					return nil, retErr
				}
				fmt.Printf("Updated Contract: %s\n", &stateContractStruct)
				return stateContractStruct, nil
			}
		}
	}
}

func AddContractItem(c router.Context) (interface{}, error) {
	contractID := c.ParamString("contract_ID")
	itemPayload := c.Param("itemPayload").(payload.Item) // Assert the chaincode parameter

	logging.CCLoggerInstance.Printf("Checking Contract Exists()\n")
	if exists, err := c.State().Exists(state.ContractState{ContractId: contractID}); err != nil {
		retErr := fmt.Errorf("Error: Exists() returned error: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else if !exists {
		retErr := fmt.Errorf("Error: Invalid delete operation, contract with ID: %s does not exist in contract state", contractID)
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	}

	if stateContract, err := c.State().Get(state.ContractState{ContractId: contractID}, &state.ContractState{}); err != nil {
		retErr := fmt.Errorf("The requested Contract does not exist: %s", err.Error())
		logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
		return nil, retErr
	} else {
		stateContractStruct := stateContract.(state.ContractState)
		stateItemsArray := stateContractStruct.Items
		stateToPayloadStruct := payload.ContractPayload{
			ContractId:   stateContractStruct.ContractId,
			ContractType: stateContractStruct.ContractType,
			Orgs:         stateContractStruct.Orgs,
			Items:        stateContractStruct.Items,
		}
		logging.CCLoggerInstance.Printf("Checking ACL rules\n")
		if owner, err := onlyContractOrgs(c); err != nil {
			retErr := fmt.Errorf("The user invoking the Contract does not belong in the ACL: %s", err.Error())
			return nil, retErr
		} else if owner != string(itemPayload.OrgId) || stateContractStruct.ContractStatus != "Accepted" {
			retErr := fmt.Errorf("The Org invoking the chaincode does not match the Item's Org")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
		logging.CCLoggerInstance.Printf("Received input: %s. Attempting to validate contract request...\n", itemPayload.String())
		if err := itemPayload.Validate(stateToPayloadStruct); err != nil {
			retErr := fmt.Errorf("Error: Validate() returned error: %s", err.Error())
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		}
		logging.CCLoggerInstance.Printf("Checking Items array length\n")
		if len(stateItemsArray) == 0 {
			retErr := fmt.Errorf("The requested Contract does not include any Items")
			logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
			return nil, retErr
		} else {
			itemFound := false
			for _, curQueriedObj := range stateItemsArray {
				if curQueriedObj.ObjectId == itemPayload.ObjectId {
					itemFound = true
				}
			}
			if itemFound {
				retErr := fmt.Errorf("The requested Item is already included into the Contract")
				logging.CCLoggerInstance.Printf("%s\n", retErr.Error())
				return nil, retErr
			} else {
				stateContractStruct.Items = append(stateContractStruct.Items, itemPayload)
				if err := c.State().Put(stateContractStruct); err != nil {
					retErr := fmt.Errorf("Error: Put() returned error: %s", err.Error())
					return nil, retErr
				}
				fmt.Printf("Updated Contract: %s\n", &stateContractStruct)
				return stateContractStruct, nil
			}
		}
	}
}
