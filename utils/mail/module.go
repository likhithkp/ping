package mail

import (
	"go.uber.org/fx"
)

var Module = fx.Module("misc-mail-ses",
	fx.Provide(
		NewMailer,
	),
)
