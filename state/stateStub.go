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

func (s StateStub) NewContract(payload payload.ContractPayload) *ContractState {
	contractState := &ContractState{
		ContractId:     payload.ContractId,
		ContractType:   payload.ContractType,
		ContractStatus: payload.ContractStatus,
		Orgs:           payload.Orgs,
		Items:          payload.Items,
		LastUpdated:    time.Now(),
		Created:        time.Now(),
	}
	if err := s.context.State().Insert(contractState); err != nil {
		return fmt.Errorf("Error: Insert() returned error: %s", err.Error())
	}
	return nil
}
