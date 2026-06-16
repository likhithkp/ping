package notification

import "go.uber.org/fx"

var Module = fx.Module("misc-notification",
	fx.Provide(NewFcmClient, NewSendNotificationUtil),
)
