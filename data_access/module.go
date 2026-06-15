package data_access

import (
	"github.com/likhithkp/ping/data_access/mongo"
	"github.com/likhithkp/ping/data_access/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("data_access",
	mongo.Module,
	repository.Module,
)
