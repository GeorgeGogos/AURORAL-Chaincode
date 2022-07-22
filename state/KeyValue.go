package state

import (
	"fmt"
)

type KeyValue struct {
	ID    string      `json:"ID"`
	Value interface{} `json:"Value"`
}

const KeyValueEntity = `KeyValue`

func (k KeyValue) Key() ([]string, error) {
	return []string{KeyValueEntity, k.ID}, nil
}

func (k KeyValue) String() string {
	return fmt.Sprintf("KeyValue(ID=%s)", k.ID)
}
