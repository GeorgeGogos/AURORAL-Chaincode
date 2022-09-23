package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/attrmgr"

	"github.com/s7techlab/cckit/router"
)

func onlyContractOrgs(c router.Context) (string, error) {
	if client, err := c.Client(); err != nil {
		retErr := fmt.Errorf("Error: Cannot retrieve Client: %s", err.Error())
		return "", retErr
	} else if clientCert, err := client.GetX509Certificate(); err != nil {
		retErr := fmt.Errorf("Error: Cannot retrieve Client: %s", err.Error())
		return "", retErr
	} else {
		mgr := attrmgr.New()
		attr, err := mgr.GetAttributesFromCert(clientCert)
		if err != nil {
			retErr := fmt.Errorf("Error: Cannot retrieve Client: %s", err.Error())
			return "", retErr
		}
		cid := attr.Attrs["cid"]
		marshaledAttr, _ := json.Marshal(cid)
		fmt.Printf("Attribute from cert is: %+v", string(marshaledAttr))
		return string(cid), nil
	}

}
