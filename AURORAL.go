package auroral

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/chaincode"
	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"

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

		Invoke(`CreateContract`, chaincode.CreateContract, param.Struct("contractPayload", &payload.ContractPayload{}))
		//Invoke(`Updatevalue`, Updatevalue, param.String(ParamKey), param.Bytes(ParamValue)).
		//Invoke(`Deletevalue`, Deletevalue, param.String(ParamKey))

	return router.NewChaincode(r)

}
