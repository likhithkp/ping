package other

import (
	"go.uber.org/fx"
)

var Module = fx.Module("utils-other",
	fx.Provide(NewUtils),
)
