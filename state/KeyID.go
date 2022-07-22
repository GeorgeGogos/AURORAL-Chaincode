package state

import (
	"fmt"
)

type KeyID struct {
	ID string `json:"ID"`
}

const KeyIDEntity = `KeyID`

func (k KeyID) Key() ([]string, error) {
	return []string{KeyIDEntity, k.ID}, nil
}

func (k KeyID) String() string {
	return fmt.Sprintf("KeyID(ID=%s)", k.ID)
}
