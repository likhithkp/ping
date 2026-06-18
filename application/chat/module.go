package chat

import (
	"github.com/likhithkp/ping/application/chat/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-chat",
	fx.Provide(
		handler.NewWsHandler,
		NewController,
	),
	fx.Invoke(RegisterChatRoutes),
)
