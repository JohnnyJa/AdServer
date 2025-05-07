package model

import (
	"github.com/google/uuid"
)

type Profile struct {
	Id        uuid.UUID
	Name      string
	Creatives map[uuid.UUID]*Creative
}
