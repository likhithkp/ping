package utils

import (
	"github.com/likhithkp/ping/utils/config"
	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/logger"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/server"
	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	config.Module,
	jwt.Module,
	logger.Module,
	other.Module,
	server.Module,
)
