package application

import (
	"github.com/likhithkp/ping/application/user"
	"go.uber.org/fx"
)

var Module = fx.Module("application",
	user.Module,
)
