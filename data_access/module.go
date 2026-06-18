package dataaccess

import (
	"github.com/likhithkp/ping/data_access/mongo"
	"github.com/likhithkp/ping/data_access/redis"
	"github.com/likhithkp/ping/data_access/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("dataaccess",
	mongo.Module,
	redis.Module,
	repository.Module,
)
