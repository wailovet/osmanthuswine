package helper

import "github.com/satori/go.uuid"

func CreateUUID() string {
	u := uuid.NewV4()
	var err error

	u1 := uuid.Must(u, err)
	return u1.String()
}
