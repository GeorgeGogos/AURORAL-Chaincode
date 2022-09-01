package payload

import (
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type Item struct {
	Enabled    null.Bool `json:"enabled"`
	Write      null.Bool `json:"write"`
	ObjectId   string    `json:"object_id"`
	UnitId     string    `json:"unit_id"`
	OrgId      string    `json:"org_id"`
	ObjectType string    `json:"object_type"`
}

func (i Item) String() string {
	return fmt.Sprintf("Enabled=%v, Write=%#v, ObjectId=%#v, UnitId=%#v, OrgId=%#v, ObjectType=%#v)",
		i.Enabled, i.Write, i.ObjectId, i.UnitId, i.OrgId, i.ObjectType)
}

func (i Item) Validate() error {
	if i.Enabled.Valid != true {
		return fmt.Errorf("Error validating Item: enabled field cannot be empty.")
	}
	if i.Write.Valid != true {
		return fmt.Errorf("Error validating Item: write field cannot be empty.")
	}
	if i.ObjectId == "" {
		return fmt.Errorf("Error validating Item: object ID cannot be an empty string.")
	}
	if i.UnitId == "" {
		return fmt.Errorf("Error validating Item: unit ID cannot be an empty string.")
	}
	if i.OrgId == "" {
		return fmt.Errorf("Error validating Item: org ID cannot be an empty string.")
	}
	if i.ObjectType == "" || (i.ObjectType != "Service" && i.ObjectType != "Device" && i.ObjectType != "Marketplace") {
		return fmt.Errorf("Error validating Item: object type cannot be an empty string.")
	}
	return nil
}
