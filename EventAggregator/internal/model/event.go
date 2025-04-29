package model

import (
	"github.com/google/uuid"
	"time"
)

type Device struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

type Meta struct {
	PlacementId int     `json:"placement_id"`
	BidPrice    float64 `json:"bid_price"`
	Currency    string  `json:"currency"`
}

type Event struct {
	RequestId   uuid.UUID `json:"request_id"`
	Timestamp   time.Time `json:"timestamp"`
	EventType   int       `json:"event_type"`
	ProfileId   int       `json:"profile_id"`
	PublisherId int       `json:"publisher_id"`
	UserId      uuid.UUID `json:"user_id"`
	Device      Device    `json:"device"`
	Meta        Meta      `json:"meta"`
}
