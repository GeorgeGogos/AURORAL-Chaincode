package chaincode

import (
	"github.com/GeorgeGogos/AURORAL-Chaincode/logging"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/router/param"
)

const ParamKey = `ParamKey`
const ParamValue = `ParamValue`

func NewCC() *router.Chaincode {
	logging.InitCCLogger()
	r := router.New(`keys`).Use(logging.SetContextMiddlewareFunc())

	r.Init(func(context router.Context) (i interface{}, e error) {
		// No implementation required with this example
		// It could be where data migration is performed, if necessary
		return nil, nil
	})

	r.
		// Read methods
		Query(`List`, List).
		// Get method has 2 params
		Query(`Get`, Get, param.String(ParamKey)).

		// Transaction methods

		Invoke(`Insertvalue`, Insertvalue, param.Bytes(ParamValue)).
		Invoke(`Updatevalue`, Updatevalue, param.String(ParamKey), param.Bytes(ParamValue)).
		Invoke(`Deletevalue`, Deletevalue, param.String(ParamKey))

	return router.NewChaincode(r)

}
