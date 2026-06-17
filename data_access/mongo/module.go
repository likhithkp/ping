package mongo

import (
	channelMongo "github.com/likhithkp/ping/data_access/mongo/channel"
	otpMongo "github.com/likhithkp/ping/data_access/mongo/otp"
	userMongo "github.com/likhithkp/ping/data_access/mongo/user"

	"go.uber.org/fx"
)

var Module = fx.Module("mongo",
	fx.Provide(
		NewClient,
		NewDatabase,
		userMongo.NewUserMongoService,
		otpMongo.NewOtpMongoService,
		channelMongo.NewChannelMongoService,
	),
)
