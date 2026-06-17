package repository

import (
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/otp"
	"github.com/likhithkp/ping/data_access/repository/user"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access-repository",
	fx.Provide(
		channel.NewChannelRepository,
		otp.NewOtpRepository,
		user.NewUserRepository,
	),
)
