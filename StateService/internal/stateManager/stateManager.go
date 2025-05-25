package stateManager

import (
	"context"
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/StateService/internal/mapper"
	"github.com/JohnnyJa/AdServer/StateService/internal/redisStorage"
	"github.com/JohnnyJa/AdServer/StateService/internal/service"
	"github.com/JohnnyJa/AdServer/StateService/internal/state"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type StateManager interface {
	service.Service
	IncrementViews(ctx context.Context, profileId uuid.UUID) error
	GetState(ctx context.Context, profileId uuid.UUID) (state.State, error)
}

type stateManager struct {
	config *app.ManagerConfig
	logger *logrus.Logger
	redis  redisStorage.RedisStorage
	client grpcClients.ProfilesLimitsClient
	cancel context.CancelFunc
}

func NewStateManager(conf *app.Config, logger *logrus.Logger, redis redisStorage.RedisStorage, client grpcClients.ProfilesLimitsClient) StateManager {
	return &stateManager{
		config: conf.ManagerConfig,
		logger: logger,
		redis:  redis,
		client: client,
	}
}

func (sm *stateManager) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	sm.cancel = cancel
	sm.startPeriodicRefresh(ctx)
	return nil
}

func (sm *stateManager) Stop(ctx context.Context) error {
	sm.cancel()
	return nil
}

func (sm *stateManager) startPeriodicRefresh(ctx context.Context) {
	ticker := time.NewTicker(sm.config.RefreshDelay)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				sm.logger.Info("Periodic refresh stopped")
				return
			case <-ticker.C:
				profiles, err := sm.client.GetProfilesWithLimits(ctx)
				if err != nil {
					sm.logger.Errorf("Periodic refresh failed, %s", err.Error())
				}
				err = sm.redis.UpdateLimits(ctx, mapper.ProtoProfileWithLimitsToMap(profiles))
				if err != nil {
					sm.logger.Errorf("Periodic refresh failed, %s", err.Error())
				}

				sm.logger.Info("Profiles refreshed at %s", time.Now().Format(time.RFC3339))
			}
		}
	}()
}

func (sm *stateManager) IncrementViews(ctx context.Context, profileId uuid.UUID) error {
	sm.logger.Infof("Incrementing views for profile %s", profileId)
	err := sm.redis.IncrementViews(ctx, profileId)
	if err != nil {
		sm.logger.Errorf("IncrementViews failed, %s", err.Error())
		return err
	}
	return nil
}

func (sm *stateManager) GetState(ctx context.Context, profileId uuid.UUID) (state.State, error) {
	status, err := sm.redis.GetProfileStatus(ctx, profileId)
	if err != nil {
		return state.Inactive, err
	}
	switch status {
	case "active":
		return state.Active, nil
	case "inactive":
		return state.Inactive, nil
	default:
		return state.Inactive, nil
	}

}
