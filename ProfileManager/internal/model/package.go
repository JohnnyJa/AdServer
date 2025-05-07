package model

import "github.com/google/uuid"

type Package map[uuid.UUID][]*Profile
