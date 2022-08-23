package test

import (
	"testing"
	"time"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	auroral "github.com/GeorgeGogos/AURORAL-Chaincode"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	testcc "github.com/s7techlab/cckit/testing"
	expectcc "github.com/s7techlab/cckit/testing/expect"
)

func TestChaincode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AURORAL Suite")
}

var (
	userCN    = "ggogos@iti.gr"
	userID, _ = GenerateCertIdentity(`SomeMSP`, userCN)
)

var _ = Describe(`Chaincode`, func() {

	//Create chaincode mock
	cc := testcc.NewMockStub(`auroral_chaincode`, auroral.NewCC())

	BeforeSuite(func() {
		// init chaincode
		expectcc.ResponseOk(cc.From(userID).Init()) // init chaincode from authority
	})

	Describe("Create", func() {

		It("Allow authority to add information about contract", func() {
			//invoke chaincode method from authority actor
			expectcc.ResponseOk(cc.From(userID).Invoke(`CreateContract`, &payload.ContractPayload{
				ContractId:     "80124570-ae01-49f5-ab04-57b7bba1c66a",
				ContractType:   "Private",
				ContractStatus: "Pending",
				Orgs:           [2]string{"3f4a58aa-d863-477a-be05-5333324b2f8d", "5f4a58aa-d863-477a-be05-4856793b2f8d"},
				Items: []payload.Item{{
					Enabled:    true,
					Write:      true,
					ObjectId:   "64c2e5d9-4829-4602-8c8d-2d26e8d00df0",
					UnitId:     "????????????????",
					OrgId:      "9c4e0166-b3f9-4f83-9192-7691b86c8b0f",
					ObjectType: "Service",
				},
					{
						Enabled:    true,
						Write:      false,
						ObjectId:   "1c44315e-981c-435d-bedb-4251f7818977",
						UnitId:     "????????????????",
						OrgId:      "3f4a58aa-d863-477a-be05-5333324b2f8d",
						ObjectType: "Device",
					}},
				LastUpdated: time.Now().Add(time.Hour * 24 * 30 * 6),
				Created:     time.Now(),
			}))

		})

	})
})
