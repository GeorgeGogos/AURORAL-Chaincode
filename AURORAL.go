package auroral

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/chaincode"
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	logging "github.com/CERTH-ITI-DLT-Lab/hlf-cc-logging"

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

		Invoke(`ProposeContract`, chaincode.ProposeContract, param.Struct("contractPayload", &payload.ContractPayload{})).
		Invoke(`AcceptContract`, chaincode.AcceptContract, param.String("contract_ID")).
		Invoke(`RejectContract`, chaincode.RejectContract, param.String("contract_ID")).
		Invoke(`DissolveContract`, chaincode.DissolveContract, param.String("contract_ID")).
		Query(`GetContractByID`, chaincode.GetContractByID, param.String("contract_ID")).
		Query(`GetContracts`, chaincode.GetContracts).
		Query(`GetContractIDs`, chaincode.GetContractIDs).
		Invoke(`UpdateContractItem`, chaincode.UpdateContractItem, param.String("contract_ID"), param.Struct("itemPayload", &payload.Item{})).
		Invoke(`DeleteContractItem`, chaincode.DeleteContractItem, param.String("contract_ID"), param.String("item_ID")).
		Invoke(`AddContractItem`, chaincode.AddContractItem, param.String("contract_ID"), param.Struct("itemPayload", &payload.Item{}))

	return router.NewChaincode(r)

}
