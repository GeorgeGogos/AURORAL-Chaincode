package test

import (
	"testing"
	"time"

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

func randomOrgs() (string, string) {
	rand.Seed(time.Now().Unix())
	org1 := uuid.New().String()
	org2 := uuid.New().String()
	return org1, org2
}

var (
	userCN             = "ggogos@iti.gr"
	org1, org2         = randomOrgs()
	randOrgs           = []string{org1, org2, "Fault Input"}
	userID, _          = GenerateCertIdentity(`SomeMSP`, userCN, org1)
	maliciousUserID, _ = GenerateCertIdentity(`SomeMSP`, userCN, randOrgs[2])
	t                  = true
	f                  = false
	Object_Type        = []string{"Service", "Device", "Marketplace"}
	Contract_Type      = []string{"Private", "Community"}
	Contract_Status    = []string{"Pending", "Approved", "Deleted"}
)

var _ = Describe(`Chaincode`, func() {

	//Create chaincode mock
	cc := testcc.NewMockStub(`auroral_chaincode`, auroral.NewCC())

	BeforeSuite(func() {
		// init chaincode
		expectcc.ResponseOk(cc.From(userID).Init()) // init chaincode from authority
	})

	Describe("Create", func() {
		It("Contract creation, expected to be succeed", func() {
			//invoke chaincode method from authority actor

			ccResponse := (cc.From(userID).Invoke(`CreateContract`, &payload.ContractPayload{
				ContractId:     uuid.New().String(),
				ContractType:   Contract_Type[rand.Intn(len(Contract_Type))],
				ContractStatus: Contract_Status[rand.Intn(len(Contract_Status))],
				Orgs:           []string{randOrgs[0], randOrgs[1]},
				Items: []payload.Item{{
					Enabled:    &t,
					Write:      &t,
					ObjectId:   uuid.New().String(),
					UnitId:     uuid.New().String(),
					OrgId:      randOrgs[rand.Intn(2)],
					ObjectType: Object_Type[rand.Intn(len(Object_Type))],
				},
					{
						Enabled:    &t,
						Write:      &f,
						ObjectId:   uuid.New().String(),
						UnitId:     uuid.New().String(),
						OrgId:      randOrgs[rand.Intn(2)],
						ObjectType: Object_Type[rand.Intn(len(Object_Type))],
					}},
			}))
			expectcc.ResponseOk(ccResponse)

		})
		It("Contract Creation with acl error, expected to fail", func() {
			ccResponse := (cc.From(maliciousUserID).Invoke(`CreateContract`, &payload.ContractPayload{
				ContractId:     uuid.New().String(),
				ContractType:   Contract_Type[rand.Intn(len(Contract_Type))],
				ContractStatus: Contract_Status[rand.Intn(len(Contract_Status))],
				Orgs:           []string{randOrgs[0], randOrgs[1]},
				Items: []payload.Item{{
					Enabled:    &t,
					Write:      &t,
					ObjectId:   uuid.New().String(),
					UnitId:     uuid.New().String(),
					OrgId:      randOrgs[rand.Intn(2)],
					ObjectType: Object_Type[rand.Intn(len(Object_Type))],
				},
					{
						Enabled:    &t,
						Write:      &f,
						ObjectId:   uuid.New().String(),
						UnitId:     uuid.New().String(),
						OrgId:      randOrgs[rand.Intn(2)],
						ObjectType: Object_Type[rand.Intn(len(Object_Type))],
					}},
			}))
			expectcc.ResponseError(ccResponse)
		})

	})
})
