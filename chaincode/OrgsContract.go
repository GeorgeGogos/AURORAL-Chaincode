package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/state"

	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"

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
	result := c.State().Insert(contractState)
	logging.CCLoggerInstance.Printf("Ledger state: %s\n", result)
	return nil, nil
}
