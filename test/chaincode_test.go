package test

import (
	"testing"
	"time"

	"gopkg.in/guregu/null.v4"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"

	"math/rand"

	auroral "github.com/GeorgeGogos/AURORAL-Chaincode"
	"github.com/google/uuid"
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
			rand.Seed(time.Now().Unix())
			expectcc.ResponseOk(cc.From(userID).Invoke(`CreateContract`, &payload.ContractPayload{
				ContractId:     uuid.New().String(),
				ContractType:   payload.contractType[rand.Intn(len(payload.contractType))],
				ContractStatus: payload.contractStatus[rand.Intn(len(payload.contractStatus))],
				Orgs:           []string{uuid.New().String(), uuid.New().String()},
				Items: []payload.Item{{
					Enabled:    null.BoolFrom(true),
					Write:      null.BoolFrom(true),
					ObjectId:   uuid.New().String(),
					UnitId:     uuid.New().String(),
					OrgId:      uuid.New().String(),
					ObjectType: payload.objectType[rand.Intn(len(payload.objectType))],
				},
					{
						Enabled:    null.BoolFrom(true),
						Write:      null.BoolFrom(false),
						ObjectId:   uuid.New().String(),
						UnitId:     uuid.New().String(),
						OrgId:      uuid.New().String(),
						ObjectType: "Device",
					}},
			}))

		})

	})
})
