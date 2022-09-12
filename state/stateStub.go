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

func (s StateStub) NewContract(payload payload.ContractPayload) error {
	txTime, _ := s.context.Time()
	contractState := &ContractState{
		ContractId:     payload.ContractId,
		ContractType:   payload.ContractType,
		ContractStatus: payload.ContractStatus,
		Orgs:           payload.Orgs,
		Items:          payload.Items,
		LastUpdated:    txTime.UTC().Format(time.RFC3339),
		Created:        txTime.UTC().Format(time.RFC3339),
	}
	fmt.Printf("%s\n", contractState.String())
	if err := s.context.State().Insert(contractState); err != nil {
		retErr := fmt.Errorf("Error: Insert() returned error: %s", err.Error())
		return retErr
	}
	return nil
}
