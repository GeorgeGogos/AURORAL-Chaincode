package main

import (
	"fmt"

	chaincode "github.com/GeorgeGogos/AURORAL-Chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	cc := chaincode.NewCC()
	if err := shim.Start(cc); err != nil {
		fmt.Printf("Error while attempting to start chaincode: %s\n", err.Error())
	}
}
