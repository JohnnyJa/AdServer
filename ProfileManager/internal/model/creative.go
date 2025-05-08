package model

import "github.com/google/uuid"

type Creative struct {
	ID           uuid.UUID
	MediaURL     string
	Width        int32
	Height       int32
	CreativeType string
}
