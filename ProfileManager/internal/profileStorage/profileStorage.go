package profileStorage

import (
	"context"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients/proto"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/model"
	"github.com/JohnnyJa/AdServer/ProfileManager/service"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ProfileStorage interface {
	service.Service
	GetPackagesByZone(zoneId uuid.UUID) []uuid.UUID
	GetProfilesByPackage(packageId uuid.UUID) []uuid.UUID
	GetProfileById(profileId uuid.UUID) *model.Profile
	GetProfilesByIds(profileIds []uuid.UUID) map[uuid.UUID]model.Profile
	GetProfilesByPackages(packageIds []uuid.UUID) map[uuid.UUID][]uuid.UUID
}

type profileStorage struct {
	config            *Config
	profilesByUUID    map[uuid.UUID]*model.Profile
	packagesByZone    map[uuid.UUID][]uuid.UUID
	profilesByPackage map[uuid.UUID][]uuid.UUID
	profileClient     gRPCClients.ProfileClient
	packageClient     gRPCClients.PackageClient
	logger            *logrus.Logger
	mu                sync.RWMutex
	cancel            context.CancelFunc
}

func New(config *Config, logger *logrus.Logger, profileClient gRPCClients.ProfileClient, packageClient gRPCClients.PackageClient) ProfileStorage {
	return &profileStorage{config: config, logger: logger, mu: sync.RWMutex{}, profileClient: profileClient, packageClient: packageClient}
}

func (s *profileStorage) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.startPeriodicRefresh(ctx, s.config.Interval)
	return nil
}

func (s *profileStorage) Stop() error {
	s.cancel()
	return nil
}

func (s *profileStorage) startPeriodicRefresh(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				s.logger.Info("Periodic refresh stopped")
				return
			case <-ticker.C:
				err := s.UpdateAll()
				if err != nil {
					s.logger.Errorf("Periodic refresh failed, %s", err.Error())
				}
				s.logger.Info("Profiles refreshed at %s", time.Now().Format(time.RFC3339))
			}
		}
	}()
}

func (s *profileStorage) UpdateAll() error {
	profiles, err := s.profileClient.GetProfiles()
	if err != nil {
		return err
	}
	packageIds := getPackageIds(profiles)

	packages, err := s.packageClient.GetPackages(packageIds)
	if err != nil {
		return err
	}

	packagesByZones := make(map[uuid.UUID][]uuid.UUID)
	for _, p := range packages {
		for _, zoneId := range p.ZoneIds {
			id := uuid.MustParse(zoneId)
			if _, exist := packagesByZones[id]; !exist {
				packagesByZones[id] = make([]uuid.UUID, 0)
			}
			packagesByZones[id] = append(packagesByZones[id], uuid.MustParse(p.Id))
		}
	}

	profilesByPackage := make(map[uuid.UUID][]uuid.UUID)
	profilesByUUID := make(map[uuid.UUID]*model.Profile)

	for _, p := range profiles {
		for _, packageId := range p.PackageIds {
			id := uuid.MustParse(packageId)
			if _, exist := profilesByPackage[id]; !exist {
				profilesByPackage[id] = make([]uuid.UUID, 0)
			}
			profilesByPackage[id] = append(profilesByPackage[id], uuid.MustParse(p.Id))
		}
		profileId := uuid.MustParse(p.Id)
		profilesByUUID[profileId] = &model.Profile{
			Id:        uuid.MustParse(p.Id),
			Name:      p.Name,
			Creatives: make(map[uuid.UUID]*model.Creative),
		}

		for _, creative := range p.Creatives {
			creativeId := uuid.MustParse(creative.Id)
			profilesByUUID[profileId].Creatives[creativeId] = &model.Creative{
				ID:                creativeId,
				MediaURL:          creative.MediaUrl,
				Width:             creative.Width,
				Height:            creative.Height,
				CreativeType:      creative.CreativeType,
				CreativeTargeting: creative.CreativeTargeting,
			}
		}
	}

	s.mu.Lock()
	s.profilesByUUID = profilesByUUID
	s.profilesByPackage = profilesByPackage
	s.packagesByZone = packagesByZones
	s.mu.Unlock()

	return nil
}

func getPackageIds(profiles []*proto.Profile) []string {
	res := make([]string, len(profiles))
	for i, profile := range profiles {
		res[i] = profile.Id
	}
	return res
}

func (s *profileStorage) GetPackagesByZone(zoneId uuid.UUID) []uuid.UUID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.packagesByZone[zoneId]
}

func (s *profileStorage) GetProfilesByPackage(packageId uuid.UUID) []uuid.UUID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.profilesByPackage[packageId]
}

func (s *profileStorage) GetProfileById(profileId uuid.UUID) *model.Profile {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.profilesByUUID[profileId]
}

func (s *profileStorage) GetProfilesByIds(profileIds []uuid.UUID) map[uuid.UUID]model.Profile {

	s.mu.RLock()
	defer s.mu.RUnlock()

	res := make(map[uuid.UUID]model.Profile)

	for _, profileId := range profileIds {
		res[profileId] = *s.profilesByUUID[profileId]
	}

	return res
}

func (s *profileStorage) GetProfilesByPackages(packageIds []uuid.UUID) map[uuid.UUID][]uuid.UUID {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make(map[uuid.UUID][]uuid.UUID)
	for _, packageId := range packageIds {
		profiles := s.profilesByPackage[packageId]
		res[packageId] = make([]uuid.UUID, 0)
		for _, profile := range profiles {
			res[packageId] = append(res[packageId], profile)
		}
	}

	return res
}
