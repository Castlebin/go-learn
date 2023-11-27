package util

import "github.com/google/uuid"

func GenUUID() string {
	uu, _ := uuid.NewUUID()
	return uu.String()
}
