package state

import (
	"fmt"
	"time"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

const ContractStateEntity = `ContractState`

type ContractState struct {
	ContractId     string         `json:"contract_id"`
	ContractType   string         `json:"contract_type"`
	ContractStatus string         `json:"contract_status"`
	Orgs           []string       `json:"orgs"`
	Items          []payload.Item `json:"items"`
	LastUpdated    time.Time      `json:"last_updated"`
	Created        time.Time      `json:"created"`
}

func (s ContractState) Key() ([]string, error) {
	return []string{ContractStateEntity, s.ContractId}, nil
}

func (s ContractState) String() string {
	return fmt.Sprintf("ContractId=%#v, ContractType=%#v, ContractStatus=%#v, Orgs=%#v, Items=%#v, LastUpdated=%s, Created=%s",
		s.ContractId, s.ContractType, s.ContractStatus, s.Orgs, s.Items, s.LastUpdated.String(), s.Created.String())
}
