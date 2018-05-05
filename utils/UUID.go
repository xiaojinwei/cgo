package utils

import "github.com/satori/go.uuid"

func RandomUUID() (uuid.UUID, error)  {
	return uuid.NewV4()
}
