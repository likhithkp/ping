package mongo

import (
	userMongo "github.com/likhithkp/ping/data_access/mongo/user"

	"go.uber.org/fx"
)

var Module = fx.Module("mongo",
	fx.Provide(
		NewClient,
		NewDatabase,
		userMongo.NewUserMongoService,
	),
)
