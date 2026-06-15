package logger

import (
	"context"

	"github.com/likhithkp/ping/utils/config"
	_const "github.com/likhithkp/ping/utils/const"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle, env *config.Env) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if env.DeploymentEnv == string(_const.Deployment_Production) {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
	return logger, nil
}
