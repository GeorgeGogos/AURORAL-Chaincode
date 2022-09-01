package payload

import (
	"fmt"
)

type ContractPayload struct {
	ContractId     string   `json:"contract_id"`
	ContractType   string   `json:"contract_type"`
	ContractStatus string   `json:"contract_status"`
	Orgs           []string `json:"orgs"`
	Items          []Item   `json:"items,omitempty"`
}

func (p ContractPayload) String() string {
	return fmt.Sprintf("ContractId=%#v, ContractType=%#v, ContractStatus=%#v, Orgs=%#v, Items=%#v",
		p.ContractId, p.ContractType, p.ContractStatus, p.Orgs, p.Items)
}

func (p ContractPayload) Validate() error {
	if p.ContractId == "" {
		return fmt.Errorf("Error validating Contract payload: contract ID cannot be an empty string.")
	}
	if p.ContractType == "" || (p.ContractType != "Private" && p.ContractType != "Community") {
		return fmt.Errorf("Error validating Contract payload: contract type cannot be an empty string.")
	}
	if p.ContractStatus == "" || (p.ContractStatus != "Pending" && p.ContractStatus != "Approved" && p.ContractStatus != "Deleted") {
		return fmt.Errorf("Error validating Contract payload: contract status cannot be an empty string.")
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
