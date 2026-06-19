package redis

import (
	"github.com/likhithkp/ping/data_access/redis/chat"
	"go.uber.org/fx"
)

var Module = fx.Module("dataaccess-redis",
	fx.Provide(
		NewRedisClient,
		chat.NewChatRedisService,
	),
)
