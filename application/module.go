package application

import (
	auth "github.com/likhithkp/ping/application/auth"
	"github.com/likhithkp/ping/application/channel"
	"github.com/likhithkp/ping/application/chat"
	user "github.com/likhithkp/ping/application/user"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	auth.Module,
	channel.Module,
	chat.Module,
	user.Module,
)
