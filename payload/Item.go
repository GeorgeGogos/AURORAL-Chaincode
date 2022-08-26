package payload

import (
	"fmt"
)

type Item struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Write      bool   `json:"write,omitempty"`
	ObjectId   string `json:"object_id,omitempty"`
	UnitId     string `json:"unit_id,omitempty"`
	OrgId      string `json:"org_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf("Enabled=%#v, Write=%#v, ObjectId=%#v, UnitId=%#v, OrgId=%#v, ObjectType=%#v)",
		i.Enabled, i.Write, i.ObjectId, i.UnitId, i.OrgId, i.ObjectType)
}
