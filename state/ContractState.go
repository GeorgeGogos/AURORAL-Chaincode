package state

import (
	"time"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

type ContractState struct {
	ContractId     string         `json:"contract_id,omitempty"`
	ContractType   string         `json:"contract_type,omitempty"`
	ContractStatus string         `json:"contract_status,omitempty"`
	Orgs           [2]string      `json:"orgs,omitempty"`
	Items          []payload.Item `json:"items,omitempty"`
	LastUpdated    time.Time      `json:"last_updated,omitempty"`
	Created        time.Time      `json:"created,omitempty"`
}

const ContractStateEntity = `ContractState`
