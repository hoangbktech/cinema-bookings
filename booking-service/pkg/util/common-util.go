package util

import (
	"github.com/satori/go.uuid"
)

func GenOrderId() (*uuid.UUID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &id, nil
}