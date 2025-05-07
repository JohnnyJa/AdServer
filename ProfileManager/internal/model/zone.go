package model

import "github.com/google/uuid"

type Zone map[uuid.UUID][]*Package
