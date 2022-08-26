package payload

import (
	"fmt"
	"time"
)

type ContractPayload struct {
	ContractId     string    `json:"contract_id,omitempty"`
	ContractType   string    `json:"contract_type,omitempty"`
	ContractStatus string    `json:"contract_status,omitempty"`
	Orgs           [2]string `json:"orgs,omitempty"`
	Items          []Item    `json:"items,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
	Created        time.Time `json:"created,omitempty"`
}

func (p ContractPayload) String() string {
	return fmt.Sprintf("ContractId=%#v, ContractType=%#v, ContractStatus=%#v, Orgs=%#v, Items=%#v, LastUpdated=%#v, Created=%#v",
		p.ContractId, p.ContractType, p.ContractStatus, p.Orgs, p.Items, p.LastUpdated.String(), p.Created.String())
}
