package main

import (
	"github.com/labstack/echo/v4"

	"Template_Echo/src/config"
	module_app "Template_Echo/src/module/app"
	"fmt"
)

func main() {
	config.LoadEnv()

	appPort := config.AppPort()
	e := echo.New()

	module_app.AppModule(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appPort)))
}
