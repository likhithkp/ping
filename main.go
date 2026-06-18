package main

import (
	"github.com/likhithkp/ping/application"
	dataaccess "github.com/likhithkp/ping/data_access"
	"github.com/likhithkp/ping/utils"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		application.Module,
		dataaccess.Module,
		utils.Module,
	).Run()
}
