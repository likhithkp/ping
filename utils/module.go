package utils

import (
	"github.com/likhithkp/ping/utils/config"
	"github.com/likhithkp/ping/utils/logger"
	"github.com/likhithkp/ping/utils/server"
	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	config.Module,
	server.Module,
	logger.Module,
)
