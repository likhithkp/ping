package repository

import (
	"github.com/likhithkp/ping/data_access/repository/channel"
	"github.com/likhithkp/ping/data_access/repository/chat"
	"github.com/likhithkp/ping/data_access/repository/otp"
	"github.com/likhithkp/ping/data_access/repository/user"
	"go.uber.org/fx"
)

var Module = fx.Module("dataaccess-repository",
	fx.Provide(
		channel.NewChannelRepository,
		otp.NewOtpRepository,
		user.NewUserRepository,
		chat.NewChatRepository,
	),
)
