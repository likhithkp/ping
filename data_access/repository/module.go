package repository

import (
	"github.com/likhithkp/ping/data_access/repository/user"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access-repository",
	fx.Provide(
		user.NewUserRepository,
	),
)
