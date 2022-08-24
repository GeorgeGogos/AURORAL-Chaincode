package payload

import (
	"fmt"
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
	Orgs           [2]string `json:"orgs,omitempty"`
	Items          []Item    `json:"items,omitempty"`
	LastUpdated    time.Time `json:"last_updated,omitempty"`
	Created        time.Time `json:"created,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf("Enabled=%v, Write=%v, ObjectId=%s, UnitId=%s, OrgId=%s, ObjectType=%s)",
		i.Enabled, i.Write, i.ObjectId, i.UnitId, i.OrgId, i.ObjectType)
}

func (p ContractPayload) String() string {
	return fmt.Sprintf("ContractId=%s, ContractType=%s, ContractStatus=%s, Orgs=%s, Items=%v, LastUpdated=%s, Created=%s)",
		p.ContractId, p.ContractType, p.ContractStatus, p.Orgs, p.Items, p.LastUpdated.String(), p.Created.String())
}
