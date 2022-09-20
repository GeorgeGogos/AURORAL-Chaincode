package chaincode

import (
	"encoding/asn1"
	"encoding/json"
	"fmt"

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
		checkID := asn1.ObjectIdentifier{1, 2, 3, 4, 5, 6, 7, 8, 1}
		for _, ext := range clientCert.Extensions {
			if ext.Id.Equal(checkID) {
				extValue := string(ext.Value)
				fmt.Printf("Organization ID: %s\n", extValue)
				type CertExt struct {
					Attr string
					Cid  string
				}
				var attr CertExt
				err := json.Unmarshal(ext.Value, &attr)
				if err != nil {

					fmt.Printf("Error geting OrgsId: %s", err.Error())
				}
				fmt.Print("Test1")
				return attr.Cid, nil
			}

		}

	}

	return "", nil
}
