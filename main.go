package main

import (
	"github.com/likhithkp/ping/application"
	"github.com/likhithkp/ping/data_access"
	"github.com/likhithkp/ping/utils"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		application.Module,
		data_access.Module,
		utils.Module,
	).Run()
}
