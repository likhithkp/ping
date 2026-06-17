package application

import (
	auth "github.com/likhithkp/ping/application/auth"
	user "github.com/likhithkp/ping/application/user"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	auth.Module,
	user.Module,
)
