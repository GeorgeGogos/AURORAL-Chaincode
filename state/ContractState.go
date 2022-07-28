package state

import (
	"time"
)

type Item struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Write      bool   `json:"write,omitempty"`
	ObjectId   string `json:"object_id,omitempty"`
	UnitId     string `json:"unit_id,omitempty"`
	OrgId      string `json:"org_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
}

type ContractPayload struct {
	ContractId     string    `json:"contract_id,omitempty"`
	ContractType   string    `json:"contract_type,omitempty"`
	ContractStatus string    `json:"contract_status,omitempty"`
	Orgs           string    `json:"orgs,omitempty"`
	Items          []Item    `json:"items,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
	Created        time.Time `json:"created,omitempty"`
}

type ContractState struct {
	ContractId     string    `json:"contract_id,omitempty"`
	ContractType   string    `json:"contract_type,omitempty"`
	ContractStatus string    `json:"contract_status,omitempty"`
	Orgs           string    `json:"orgs,omitempty"`
	Items          []Item    `json:"items,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
	Created        time.Time `json:"created,omitempty"`
}

const ContractStateEntity = `ContractState`
