package model

import "github.com/google/uuid"

type Event struct {
	Timestamp   string    `json:"timestamp"`
	RequestID   string    `json:"request_id"`
	EventType   int       `json:"event_type"`
	ProfileID   uuid.UUID `json:"profile_id"`
	PublisherID int       `json:"publisher_id"`
	UserID      string    `json:"user_id"`
	IP          string    `json:"ip"`
	UserAgent   string    `json:"user_agent"`
	PlacementID int       `json:"placement_id"`
	BidPrice    float64   `json:"bid_price"`
	Currency    string    `json:"currency"`
}
