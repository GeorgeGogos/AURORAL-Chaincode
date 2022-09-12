package payload

import (
	"fmt"
)

type Item struct {
	Enabled    *bool  `json:"enabled"`
	Write      *bool  `json:"write"`
	ObjectId   string `json:"object_id"`
	UnitId     string `json:"unit_id"`
	OrgId      string `json:"org_id"`
	ObjectType string `json:"object_type"`
}

func (i Item) String() string {
	return fmt.Sprintf("Enabled=%b, Write=%b, ObjectId=%s, UnitId=%s, OrgId=%s, ObjectType=%s)",
		i.Enabled, i.Write, i.ObjectId, i.UnitId, i.OrgId, i.ObjectType)
}

func (i Item) Validate(p ContractPayload) error {
	if i.Enabled == nil {
		return fmt.Errorf("Error validating Item: Enabled field cannot be empty.")
	}
	if i.Write == nil {
		return fmt.Errorf("Error validating Item: Write field cannot be empty.")
	}
	if i.ObjectId == "" {
		return fmt.Errorf("Error validating Item: ObjectID cannot be an empty string.")
	}
	if i.UnitId == "" {
		return fmt.Errorf("Error validating Item: UniID cannot be an empty string.")
	}
	if i.OrgId == "" || (i.OrgId != p.Orgs[0] && i.OrgId != p.Orgs[1]) {
		return fmt.Errorf("Error validating Item: OrgID does not match the Organisations IDs.")
	}
	if i.ObjectType == "" || (i.ObjectType != "Service" && i.ObjectType != "Device" && i.ObjectType != "Marketplace") {
		return fmt.Errorf("Error validating Item: ObjectType cannot be an empty string.")
	}
	return nil
}
