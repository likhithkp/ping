package mail

import (
	"github.com/likhithkp/ping/utils/mail/ses"

	"go.uber.org/fx"
)

var Module = fx.Module("misc-mail",
	ses.Module,
)
