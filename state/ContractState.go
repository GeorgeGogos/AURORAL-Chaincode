package state

import (
	"fmt"
	"time"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

const ContractStateEntity = `ContractState`

type ContractState struct {
	ContractId     string         `json:"contract_id,omitempty"`
	ContractType   string         `json:"contract_type,omitempty"`
	ContractStatus string         `json:"contract_status,omitempty"`
	Orgs           [2]string      `json:"orgs,omitempty"`
	Items          []payload.Item `json:"items,omitempty"`
	LastUpdated    time.Time      `json:"last_updated,omitempty"`
	Created        time.Time      `json:"created,omitempty"`
}

func (s ContractState) Key() ([]string, error) {
	return []string{ContractStateEntity, s.ContractId}, nil
}

func (s ContractState) String() string {
	return fmt.Sprintf("ContractId=%#v, ContractType=%#v, ContractStatus=%#v, Orgs=%#v, Items=%#v, LastUpdated=%#v, Created=%#v",
		s.ContractId, s.ContractType, s.ContractStatus, s.Orgs, s.Items, s.LastUpdated, s.Created)
}
