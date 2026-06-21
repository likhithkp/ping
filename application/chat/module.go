package chat

import (
	"github.com/likhithkp/ping/application/chat/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("application-chat",
	fx.Provide(
		handler.NewAckMessageHandler,
		handler.NewFetchOfflineMesagesHandler,
		handler.NewReadMessageHandler,
		handler.NewSendWithRetryHandler,
		handler.NewSendPongHandler,
		handler.NewWriteMessageHandler,
		handler.NewWsConnectHandler,
		NewController,
	),
	fx.Invoke(RegisterChatRoutes),
)
