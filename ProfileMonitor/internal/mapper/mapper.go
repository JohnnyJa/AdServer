package mapper

import (
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
				Id:         row.ProfileID,
				Name:       row.ProfileName,
				Creatives:  map[uuid.UUID]*model.Creative{},
				PackageIDs: row.PackageIDs,
			}
			profiles[row.ProfileID] = profile
		}

		creative, exists := profile.Creatives[row.CreativeID]
		if !exists {
			creative = &model.Creative{
				ID:                row.CreativeID,
				MediaURL:          row.MediaURL,
				Width:             row.Width,
				Height:            row.Height,
				CreativeType:      row.CreativeType,
				CreativeTargeting: map[string]string{},
			}
			profile.Creatives[row.CreativeID] = creative
		}

		if row.TargetingKey != nil && row.TargetingValue != nil {
			creative.CreativeTargeting[*row.TargetingKey] = *row.TargetingValue
		}

	}
	return profiles, nil
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
			Targeting:    creative.CreativeTargeting,
		})
	}

	return k
}
