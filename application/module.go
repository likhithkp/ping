package application

import (
	user "github.com/likhithkp/ping/application/auth"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	user.Module,
)
