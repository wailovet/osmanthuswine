package helper

import "github.com/google/uuid"

func CreateUUID() string {
	uid := uuid.Must(uuid.NewUUID())
	return uid.String()
}
