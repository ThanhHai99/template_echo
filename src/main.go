package main

import (
	"Template_Echo/src/config"
	"Template_Echo/src/module/app"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	appPort := config.AppPort()
	e := echo.New()

	app.AppModule(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appPort)))
}
