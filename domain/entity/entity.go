package entity

import "github.com/google/uuid"

// ID entity ID
type ID = uuid.UUID

// NewID creates a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

// StringToID converts a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
