package state

import (
	"encoding/json"
	"fmt"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

const ContractStateEntity = `ContractState`

type ContractState struct {
	ContractId     string         `json:"contract_id"`
	ContractType   string         `json:"contract_type"`
	ContractStatus string         `json:"contract_status"`
	Orgs           []string       `json:"orgs"`
	Items          []payload.Item `json:"items,omitempty"`
	LastUpdated    string         `json:"last_updated"`
	Created        string         `json:"created"`
}

func (s ContractState) Key() ([]string, error) {
	return []string{ContractStateEntity, s.ContractId}, nil
}

func (s ContractState) String() string {
	marshaledItem, _ := json.Marshal(s.Items)
	return fmt.Sprintf("ContractId=%s, ContractType=%s, ContractStatus=%s, Orgs=%s, Items=%s, LastUpdated=%s, Created=%s",
		s.ContractId, s.ContractType, s.ContractStatus, s.Orgs, marshaledItem, s.LastUpdated, s.Created)
}
