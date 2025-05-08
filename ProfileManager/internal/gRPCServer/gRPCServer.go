package gRPCServer

import (
	"context"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCServer/proto"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/profileStorage"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedProfilesByZoneServiceServer
	storage profileStorage.ProfileStorage
	logger  *logrus.Logger
}

func Register(gRPC *grpc.Server, storage profileStorage.ProfileStorage, logger *logrus.Logger) {
	proto.RegisterProfilesByZoneServiceServer(gRPC, &server{storage: storage, logger: logger})
}

func (s *server) GetProfilesByZone(ctx context.Context, req *proto.GetProfileByZoneRequest) (*proto.GetProfilesByZoneResponse, error) {
	zoneUUID, err := uuid.Parse(req.ZoneId)
	if err != nil {
		return nil, err
	}

	packages := s.storage.GetPackagesByZone(zoneUUID)

	profilesByPackages := s.storage.GetProfilesByPackages(packages)
	profileIds := make([]uuid.UUID, 0)

	responseProfilesByPackages := make(map[string]*proto.ProfileIds)
	for packageId, profiles := range profilesByPackages {
		parsedProfiles := make([]string, 0)
		for _, profile := range profiles {
			profileIds = append(profileIds, profile)
			parsedProfiles = append(parsedProfiles, profile.String())
		}
		responseProfilesByPackages[packageId.String()] = &proto.ProfileIds{
			Ids: parsedProfiles,
		}
	}

	profilesById := s.storage.GetProfilesByIds(profileIds)
	responseProfilesByIds := make(map[string]*proto.Profile)
	for id, profile := range profilesById {
		idString := id.String()

		responseProfilesByIds[idString] = &proto.Profile{
			Id:                idString,
			Name:              profile.Name,
			BidPrice:          profile.BidPrice,
			Creatives:         make(map[string]*proto.Creative),
			ProfilesTargeting: profile.ProfileTargeting,
		}
		for creativeId, creative := range profile.Creatives {
			responseProfilesByIds[idString].Creatives[creativeId.String()] = &proto.Creative{
				Id:           creative.ID.String(),
				MediaURL:     creative.MediaURL,
				Width:        creative.Width,
				Height:       creative.Height,
				CreativeType: creative.CreativeType,
			}
		}
	}

	resp := &proto.GetProfilesByZoneResponse{
		ProfilesByPackage: responseProfilesByPackages,
		ProfilesByUUID:    responseProfilesByIds,
	}
	return resp, nil
}
