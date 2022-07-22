package test

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-protos-go/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testcc "github.com/s7techlab/cckit/testing"
	expectcc "github.com/s7techlab/cckit/testing/expect"
	chaincode "iti-gitlab.iti.gr/hyperledger-fabric-tools/key-value-chaincode"
	"iti-gitlab.iti.gr/hyperledger-fabric-tools/key-value-chaincode/state"
	"iti-gitlab.iti.gr/hyperledger-fabric-tools/key-value-chaincode/testdata"
)

func TestKeyValue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Key Value Suite")
}

var _ = Describe(`KeyValue`, func() {

	Describe("Key Value lifecycle", func() {
		keyChaincode := testcc.NewMockStub(`keys`, chaincode.NewCC())
		userCN := "govadmin1@example.com"
		userID, _ := GenerateCertIdentity(testdata.DefaultMSP, userCN)
		var ccResponse peer.Response
		var keyinserted state.KeyValue
		var updatedvalue state.KeyValue
		BeforeSuite(func() {
			expectcc.ResponseOk(keyChaincode.From(userID).Init())
		})

		It("Allow user to insert new key", func() {
			//invoke chaincode method from authority actor
			ccResponse = keyChaincode.From(userID).Invoke(`Insertvalue`, []byte("qwertyuio"))
			expectcc.ResponseOk(ccResponse)
			keyinserted = expectcc.PayloadIs(ccResponse, &state.KeyValue{}).(state.KeyValue)
			fmt.Printf("to value einai:%s", string(keyinserted.Value))
			ccResponse = keyChaincode.From(userID).Query("Get", keyinserted.ID)
			expectcc.ResponseOk(ccResponse)
			ccResponse := keyChaincode.From(userID).Query("List")
			expectcc.ResponseOk(ccResponse)
			KeyArray := expectcc.PayloadIs(ccResponse, &[]state.KeyValue{}).([]state.KeyValue)
			Expect(KeyArray).To(HaveLen(1))

			//expectcc.ResponseOk(keyChaincode.From(userID).Invoke(`Insertvalue`, []byte("qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[sdfghjkl;xcvbnm,.wertyuioasdfghjklzxcvbnm,.")))

		})

		It("Allow user to get key", func() {

			ccResponse = keyChaincode.From(userID).Query("Get", keyinserted.ID)
			expectcc.ResponseOk(ccResponse)

			// queryResponse := keyChaincode.From(userID).Query("Get","sdsd")
			// keyvalue := expectcc.PayloadIs(queryResponse, &state.KeyValue{}).(state.KeyValue)
			// fmt.Printf(keyvalue.ID)
			// Expect(keyvalue.Value).To(Equal([]byte("qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[sdfghjkl;xcvbnm,.wertyuioasdfghjklzxcvbnm,.")))
		})

		It("Allow user to get a list of keys", func() {
			queryResponse := keyChaincode.From(userID).Query("List")
			expectcc.ResponseOk(queryResponse)
			KeyArray := expectcc.PayloadIs(queryResponse, &[]state.KeyValue{}).([]state.KeyValue)
			Expect(KeyArray).To(HaveLen(1))
		})

		It("Allow user to Update key", func() {
			//invoke chaincode method from authority actor
			kid := keyinserted.ID
			ccResponse = keyChaincode.From(userID).Invoke(`Updatevalue`, kid, []byte("lolsasasasassasasa"))
			expectcc.ResponseOk(ccResponse)
			ccResponse = keyChaincode.From(userID).Query("Get", kid)
			updatedvalue = expectcc.PayloadIs(ccResponse, &state.KeyValue{}).(state.KeyValue)
			fmt.Printf("to value einai:%s", string(updatedvalue.Value))
			fmt.Printf("Successfully updated")

		})

		It("Allow user to delete key", func() {

			//invoke chaincode method from authority actor

			expectcc.ResponseOk(keyChaincode.From(userID).Invoke(`Deletevalue`, keyinserted.ID))
			queryResponse := keyChaincode.From(userID).Query("List")
			expectcc.ResponseOk(queryResponse)
			KeyArray := expectcc.PayloadIs(queryResponse, &[]state.KeyValue{}).([]state.KeyValue)
			Expect(KeyArray).To(HaveLen(0))
		})
	})
})
