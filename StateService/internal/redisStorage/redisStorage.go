package redisStorage

import (
	"context"
	"fmt"
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/JohnnyJa/AdServer/StateService/internal/service"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisStorage interface {
	service.Service
	IncrementViews(ctx context.Context, profileId uuid.UUID) error
	UpdateLimits(ctx context.Context, limits map[uuid.UUID]int) error
	GetProfileStatus(ctx context.Context, profileId uuid.UUID) (string, error)
}

type redisStorage struct {
	config *app.RedisConfig
	logger *logrus.Logger
	client *redis.Client
}

func NewRedisStorage(cfg *app.Config, logger *logrus.Logger) RedisStorage {
	return &redisStorage{
		config: cfg.RedisConfig,
		logger: logger,
	}
}

func (r *redisStorage) Start(ctx context.Context) error {
	opt, err := redis.ParseURL(r.config.ConnectionString)
	if err != nil {
		r.logger.Fatal("Cannot parse connection string to Redis: ", err)
		return err
	}

	client := redis.NewClient(opt)

	err = client.Ping(ctx).Err()
	if err != nil {
		r.logger.Fatal("Cannot connect to Redis: ", err)
		return err
	}

	r.client = client

	return nil
}

func (r *redisStorage) Stop(ctx context.Context) error {
	err := r.client.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisStorage) UpdateLimits(ctx context.Context, limits map[uuid.UUID]int) error {
	pipe := r.client.Pipeline()

	script := `
local newLimit = tonumber(ARGV[1])
local views = tonumber(redis.call("HGET", KEYS[1], "views") or "0")

redis.call("HSET", KEYS[1], "limit", newLimit)

if views >= newLimit then
    redis.call("HSET", KEYS[1], "status", "stopped")
else
    redis.call("HSET", KEYS[1], "status", "active")
end

return 1
`

	for id, limit := range limits {
		key := fmt.Sprintf("profile:%s", id.String())

		pipe.Eval(ctx, script, []string{key}, limit)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		r.logger.Fatal("Limits wasn't updated: ", err)
		return err
	}

	r.logger.Infof("Limits updated")

	return nil
}

func (r *redisStorage) IncrementViews(ctx context.Context, profileId uuid.UUID) error {
	key := fmt.Sprintf("profile:%s", profileId.String())

	lua := `
    local key = KEYS[1]
    local new_status = "stopped"
    local views = tonumber(redis.call("HINCRBY", key, "views", 1))
    local limit = tonumber(redis.call("HGET", key, "limit") or "0")
    local status_updated = 0
    if views >= limit then
        redis.call("HSET", key, "status", new_status)
        status_updated = 1
    end
    return status_updated
    `

	res, err := r.client.Eval(ctx, lua, []string{key}).Result()
	if err != nil {
		return err
	}

	updated := res.(int64)
	if updated == 1 {
		r.logger.Infof("Profile %s was stopped", profileId.String())
	}

	return nil
}

func (r *redisStorage) GetProfileStatus(ctx context.Context, profileId uuid.UUID) (string, error) {
	key := fmt.Sprintf("profile:%s", profileId.String())

	status, err := r.client.HGet(ctx, key, "status").Result()
	if err != nil {
		return "stopped", err //if we can't get profile status better to think it's stopped
	} else {
		return status, nil
	}
}
