package payload

import (
	"encoding/json"
	"fmt"
)

type ContractPayload struct {
	ContractId   string   `json:"contract_id"`
	ContractType string   `json:"contract_type"`
	Orgs         []string `json:"orgs"`
	Items        []Item   `json:"items,omitempty"`
}
type ContractPayloadAllias ContractPayload

func (p ContractPayload) String() string {
	marshaledItem, _ := json.Marshal(p.Items)
	return fmt.Sprintf("ContractPayload (ContractId=%s, ContractType=%s, Orgs=%s, Items=%s)",
		p.ContractId, p.ContractType, p.Orgs, string(marshaledItem))
}

func (p ContractPayload) Validate() error {
	if p.ContractId == "" {
		return fmt.Errorf("Error validating Contract payload: ContractID cannot be an empty string.")
	}
	if p.ContractType == "" || (p.ContractType != "Private" && p.ContractType != "Community") {
		return fmt.Errorf("Error validating Contract payload: ContractType cannot be an empty string.")
	}
	if len(p.Orgs) != 2 {
		return fmt.Errorf("Error validating Contract payload: contracted Orgs must be two (2).")
	}
	for i := 0; i < len(p.Orgs); i++ {
		if p.Orgs[i] == "" {
			return fmt.Errorf("Error validating Contract payload: contracted Orgs cannot be an empty string.")
		}
	}
	return nil
}
