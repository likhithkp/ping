package channel

import (
	"github.com/likhithkp/ping/application/channel/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-channel",
	fx.Provide(
		handler.NewCreateChannelHandler,
		handler.NewGetChannelListHandler,
		handler.NewGetChannelDetailsHandler,
		handler.NewDeleteChannelHandler,
		NewController,
	),
	fx.Invoke(RegisterChannelRoutes),
)
