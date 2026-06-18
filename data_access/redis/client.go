package redis

import (
	"context"

	"github.com/likhithkp/ping/utils/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewRedisClient(lc fx.Lifecycle, env *config.Env, logger *zap.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddress,
		Password: env.RedisPassword,
		DB:       0,
	})

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing Redis connection...")
			return client.Close()
		},
	})

	logger.Info("Redis client initialized successfully")
	return client, nil
}
