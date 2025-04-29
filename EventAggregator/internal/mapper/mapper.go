package mapper

import (
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/model"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/druid-repo"
)

func MapToDruidEvent(e model.Event) druid.Event {
	return druid.Event{
		Timestamp:   e.Timestamp.String(),
		RequestID:   e.RequestId.String(),
		EventType:   e.EventType,
		ProfileID:   e.ProfileId,
		PublisherID: e.PublisherId,
		UserID:      e.UserId.String(),
		IP:          e.Device.Ip,
		UserAgent:   e.Device.UserAgent,
		PlacementID: e.Meta.PlacementId,
		BidPrice:    e.Meta.BidPrice,
		Currency:    e.Meta.Currency,
	}
}
