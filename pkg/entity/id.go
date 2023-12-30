package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(text string) (ID, error) {
	id, err := uuid.Parse(text)
	return ID(id), err
}
