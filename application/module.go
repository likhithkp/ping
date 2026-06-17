package application

import (
	auth "github.com/likhithkp/ping/application/auth"
	"github.com/likhithkp/ping/application/channel"
	user "github.com/likhithkp/ping/application/user"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	auth.Module,
	channel.Module,
	user.Module,
)
