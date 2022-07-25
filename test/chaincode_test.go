package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GeorgeGogos/AURORAL-Chaincode/chaincode"
	"github.com/GeorgeGogos/AURORAL-Chaincode/state"
	"github.com/GeorgeGogos/AURORAL-Chaincode/testdata"
	"github.com/hyperledger/fabric-protos-go/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testcc "github.com/s7techlab/cckit/testing"
	expectcc "github.com/s7techlab/cckit/testing/expect"
)

func TestKeyValue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chaincode Suite")
}

var _ = Describe(`AuroralChaincode`, func() {

	Describe("Chaincode lifecycle", func() {
		auroralChaincode := testcc.NewMockStub(`auroral_chaincode`, chaincode.NewCC())
		userCN := "govadmin1@example.com"
		userID, _ := GenerateCertIdentity(testdata.DefaultMSP, userCN)
		var ccResponse peer.Response
		BeforeSuite(func() {
			expectcc.ResponseOk(auroralChaincode.From(userID).Init())
		})

		It("Allow Orgs to create a new contract", func() {
			//invoke chaincode method from authority actor
			ccResponse = auroralChaincode.Invoke(`createcontract`, &state.ContractPayload{
				ContractId:     "80124570-ae01-49f5-ab04-57b7bba1c66a",
				ContractType:   "Private",
				ContractStatus: "Pending",
				Orgs:           []string{"3f4a58aa-d863-477a-be05-5333324b2f8d", "9c4e0166-b3f9-4f83-9192-7691b86c8b0f"},
				Items: state.Item{
					Enabled:    true,
					Write:      true,
					ObjectId:   "1c44315e-981c-435d-bedb-4251f7818977",
					UnitId:     "????????????????",
					OrgId:      "3f4a58aa-d863-477a-be05-5333324b2f8d",
					ObjectType: "Device",
				},
				LastUpdated: time.Now().Add(time.Hour * 24 * 30 * 6),
				Created:     time.Now(),
			})
			expectcc.ResponseOk(ccResponse)
			payloadinserted := expectcc.PayloadIs(ccResponse, &state.ContractState{}).(state.ContractState)
			fmt.Printf("The Contract State is:%s", string(payloadinserted))

			//expectcc.ResponseOk(keyChaincode.From(userID).Invoke(`Insertvalue`, []byte("qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[sdfghjkl;xcvbnm,.wertyuioasdfghjklzxcvbnm,.")))

		})

	})
})
