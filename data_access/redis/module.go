package redis

import (
	"github.com/likhithkp/ping/data_access/redis/chat"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access-redis",
	fx.Provide(
		NewRedisClient,
		chat.NewUrlRedisService,
	),
)
