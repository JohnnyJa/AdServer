package mapper

import (
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients/proto"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/model"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/requests"
	"github.com/google/uuid"
)

func ProtoToProfilesByUUID(data map[string]*proto.Profile) map[uuid.UUID]*model.Profile {
	result := make(map[uuid.UUID]*model.Profile)
	for profileId, profile := range data {
		profileIdUUID := uuid.MustParse(profileId)
		result[profileIdUUID] = &model.Profile{
			Id:               profileIdUUID,
			Name:             profile.Name,
			BidPrice:         profile.BidPrice,
			Creatives:        map[uuid.UUID]*model.Creative{},
			ProfileTargeting: profile.ProfilesTargeting,
		}

		for creativeID, creative := range profile.Creatives {
			result[profileIdUUID].Creatives[uuid.MustParse(creativeID)] = &model.Creative{
				ID:           uuid.MustParse(creativeID),
				MediaURL:     creative.MediaURL,
				Width:        int(creative.Width),
				Height:       int(creative.Height),
				CreativeType: creative.CreativeType,
			}
		}
	}
	return result
}

func ProtoToProfilesByPackages(data map[string]*proto.ProfileIds) map[uuid.UUID][]uuid.UUID {
	result := make(map[uuid.UUID][]uuid.UUID)

	for packageId, profileIds := range data {
		var profiles []uuid.UUID
		for _, id := range profileIds.Ids {
			profiles = append(profiles, uuid.MustParse(id))
		}

		result[uuid.MustParse(packageId)] = profiles
	}
	return result
}

func NewBidResponse(bidRequest requests.BidRequest, winners map[uuid.UUID]*model.ProfileWithMatchedCreative) *requests.BidResponse {
	bidResponse := &requests.BidResponse{}
	bidResponse.ID = bidRequest.ID
	bidResponse.BidID = uuid.New().String()
	bidResponse.Cur = "USD"
	bidResponse.SeatBid = make([]requests.SeatBid, 0)
	seatBid := requests.SeatBid{
		Bid:   make([]requests.Bid, 0),
		Seat:  "Really cool bidder",
		Group: 0,
	}
	bidResponse.SeatBid = append(bidResponse.SeatBid, seatBid)

	for _, imp := range bidRequest.Imp {
		impUUID := uuid.MustParse(imp.ID)
		bid := requests.Bid{}
		bid.ID = uuid.New().String()
		bid.ImpID = imp.ID
		bid.Price = winners[impUUID].BidPrice
		bid.AdID = winners[impUUID].Id.String()
		bid.NURL = "Some NURL"
		bid.Adm = winners[impUUID].Creative.MediaURL
		bid.CrID = winners[impUUID].Creative.ID.String()
		bid.W = winners[impUUID].Creative.Width
		bid.H = winners[impUUID].Creative.Height
		bid.DealID = ""

		bidResponse.SeatBid[0].Bid = append(bidResponse.SeatBid[0].Bid, bid)
	}
	return bidResponse
}
