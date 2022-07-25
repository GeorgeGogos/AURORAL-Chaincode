package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"
	"github.com/GeorgeGogos/AURORAL-Chaincode/state"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/router/param"
)

func NewCC() *router.Chaincode {
	logging.InitCCLogger()
	r := router.New(`auroral_chaincode`).Use(logging.SetContextMiddlewareFunc())

	r.Init(func(context router.Context) (i interface{}, e error) {
		// No implementation required with this example
		// It could be where data migration is performed, if necessary
		return nil, nil
	})

	r.
		// Read methods
		//Query(`List`, List).
		// Get method has 2 params
		//Query(`Get`, Get, param.String(ParamKey)).

		// Transaction methods

		Invoke(`createcontract`, CreateContract, param.Struct("contractPayload", &state.ContractPayload{}))
		//Invoke(`Updatevalue`, Updatevalue, param.String(ParamKey), param.Bytes(ParamValue)).
		//Invoke(`Deletevalue`, Deletevalue, param.String(ParamKey))

	return router.NewChaincode(r)

}

func CreateContract(c router.Context) (interface{}, error) {
	var (
		contractPayload = c.Param("contractPayload").(state.ContractPayload) // Assert the chaincode parameter
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

	return contractState, c.State().Insert(contractState)
}
