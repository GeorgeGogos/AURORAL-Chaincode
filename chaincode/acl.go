package chaincode

import (
	"fmt"

	"github.com/s7techlab/cckit/router"
)

func onlyContractOrgs(c router.Context) string {
	client, _ := c.Client()
	clientCert, _ := client.GetX509Certificate()
	serialNumber := clientCert.Subject.SerialNumber
	fmt.Printf("Organization ID: %s\n", serialNumber)
	return serialNumber
}
