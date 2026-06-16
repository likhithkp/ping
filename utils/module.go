package utils

import (
	"github.com/likhithkp/ping/utils/aws"
	"github.com/likhithkp/ping/utils/config"
	"github.com/likhithkp/ping/utils/jwt"
	"github.com/likhithkp/ping/utils/logger"
	"github.com/likhithkp/ping/utils/mail"
	"github.com/likhithkp/ping/utils/middleware"
	"github.com/likhithkp/ping/utils/notification"
	"github.com/likhithkp/ping/utils/other"
	"github.com/likhithkp/ping/utils/server"
	"github.com/likhithkp/ping/utils/storage"
	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	aws.Module,
	config.Module,
	jwt.Module,
	logger.Module,
	mail.Module,
	middleware.Module,
	notification.Module,
	other.Module,
	server.Module,
	storage.Module,
)
