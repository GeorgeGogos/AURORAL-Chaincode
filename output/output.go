package output

import (
	"encoding/json"
	"fmt"

	"github.com/GeorgeGogos/AURORAL-Chaincode/payload"
)

type OutputContract struct {
	ContractId     string         `json:"contract_id"`
	ContractType   string         `json:"contract_type"`
	ContractStatus string         `json:"contract_status"`
	Orgs           []string       `json:"orgs"`
	Items          []payload.Item `json:"items,omitempty"`
	LastUpdated    string         `json:"last_updated"`
	Created        string         `json:"created"`
}

func (s OutputContract) String() string {
	marshaledItem, _ := json.Marshal(s.Items)
	return fmt.Sprintf("OutputContract (ContractId=%s, ContractType=%s, ContractStatus=%s, Orgs=%s, Items=%s, LastUpdated=%s, Created=%s)",
		s.ContractId, s.ContractType, s.ContractStatus, s.Orgs, string(marshaledItem), s.LastUpdated, s.Created)
}
