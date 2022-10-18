package state

import (
	"fmt"
	"time"

	"github.com/s7techlab/cckit/router"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

type StateStub struct {
	context router.Context
}

func NewStateStub(c router.Context) *StateStub {
	return &StateStub{
		context: c,
	}
}

func (s StateStub) NewContract(payload payload.ContractPayload, owner string) error {
	if txTime, err := s.context.Time(); err != nil {
		retErr := fmt.Errorf("Error retrieving Transaction's Proposal Time: %s", err.Error())
		return retErr
	} else {

		contract_State := &ContractState{
			InvokingOrg:    owner,
			ContractId:     payload.ContractId,
			ContractType:   payload.ContractType,
			ContractStatus: "Pending",
			Orgs:           payload.Orgs,
			Items:          payload.Items,
			LastUpdated:    txTime.UTC().Format(time.RFC3339),
			Created:        txTime.UTC().Format(time.RFC3339),
		}
		fmt.Printf("%s\n", contract_State.String())
		if err := s.context.State().Insert(contract_State); err != nil {
			retErr := fmt.Errorf("Error: Insert() returned error: %s", err.Error())
			return retErr
		}
	}

	return nil
}
