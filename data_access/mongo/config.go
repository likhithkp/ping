package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/likhithkp/ping/utils/config"
	"github.com/likhithkp/ping/utils/ctx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewClient(lc fx.Lifecycle, env *config.Env, logger *zap.Logger) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx.Background, 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("%v%v", env.MongoUri, env.Database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error("MongoDB connection failed", zap.Error(err))
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		logger.Error("MongoDB ping failed", zap.Error(err))
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})

	logger.Info("MongoDB client initialized successfully")
	return client, nil
}

func NewDatabase(client *mongo.Client, env *config.Env) (*mongo.Database, error) {
	return client.Database(env.Database), nil
}
