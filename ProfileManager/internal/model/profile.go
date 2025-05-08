package model

import (
	"github.com/google/uuid"
)

type Profile struct {
	Id               uuid.UUID
	Name             string
	BidPrice         float32
	Creatives        map[uuid.UUID]*Creative
	ProfileTargeting map[string]string
}
