package utils

import "github.com/satori/go.uuid"

func RandomUUID() (uuid.UUID)  {
	return uuid.NewV4()
}
