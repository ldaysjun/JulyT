package julyUuid

import "github.com/google/uuid"

func GenerateUuid() string{
	uuid,_:= uuid.NewUUID()
	return uuid.String()
}

