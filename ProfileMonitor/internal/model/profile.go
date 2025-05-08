package model

import (
	"github.com/google/uuid"
)

type Profile struct {
	Id               uuid.UUID
	Name             string
	BidPrice         float64
	Creatives        map[uuid.UUID]*Creative
	PackageIDs       []uuid.UUID
	ProfileTargeting map[string]string
}
