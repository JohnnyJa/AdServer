package mapper

import (
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/gRPC/proto"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/kafka"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/model"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/repository"
	"github.com/google/uuid"
)

func ToProfiles(rows []repository.ProfileRow) (map[uuid.UUID]*model.Profile, error) {
	profiles := make(map[uuid.UUID]*model.Profile)
	for _, row := range rows {

		profile, exists := profiles[row.ProfileID]
		if !exists {
			profile = &model.Profile{
				Id:               row.ProfileID,
				Name:             row.ProfileName,
				BidPrice:         row.BidPrice,
				Creatives:        map[uuid.UUID]*model.Creative{},
				PackageIDs:       row.PackageIDs,
				ProfileTargeting: map[string]string{},
			}
			profiles[row.ProfileID] = profile
		}

		creative, exists := profile.Creatives[row.CreativeID]
		if !exists {
			creative = &model.Creative{
				ID:           row.CreativeID,
				MediaURL:     row.MediaURL,
				Width:        row.Width,
				Height:       row.Height,
				CreativeType: row.CreativeType,
			}
			profile.Creatives[row.CreativeID] = creative
		}

		if row.TargetingKey != nil && row.TargetingValue != nil {
			profile.ProfileTargeting[*row.TargetingKey] = *row.TargetingValue
		}

	}
	return profiles, nil
}

func ToGrpcProfiles(p map[uuid.UUID]*model.Profile) *proto.GetProfilesResponse {
	resp := &proto.GetProfilesResponse{}
	for _, p := range p {
		var packageIds []string
		for _, pkg := range p.PackageIDs {
			packageIds = append(packageIds, pkg.String())
		}

		creatives := make(map[string]*proto.Creative)
		for _, creative := range p.Creatives {
			creatives[creative.ID.String()] = &proto.Creative{
				Id:           creative.ID.String(),
				MediaUrl:     creative.MediaURL,
				Width:        int32(creative.Width),
				Height:       int32(creative.Height),
				CreativeType: creative.CreativeType,
			}
		}

		resp.Profiles = append(resp.Profiles, &proto.Profile{
			Id:               p.Id.String(),
			Name:             p.Name,
			BidPrice:         float32(p.BidPrice),
			PackageIds:       packageIds,
			Creatives:        creatives,
			ProfileTargeting: p.ProfileTargeting,
		})
	}
	return resp
}

func ToKafkaProfile(p model.Profile) kafka.ProfileForKafka {
	k := kafka.ProfileForKafka{
		ID:        p.Id.String(),
		Name:      p.Name,
		Creatives: make([]kafka.CreativeForKafka, 0),
	}

	for _, creative := range p.Creatives {
		k.Creatives = append(k.Creatives, kafka.CreativeForKafka{
			ID:           creative.ID.String(),
			MediaURL:     creative.MediaURL,
			Width:        creative.Width,
			Height:       creative.Height,
			CreativeType: creative.CreativeType,
		})
	}

	return k
}
