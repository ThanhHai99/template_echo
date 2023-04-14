package main

import (
	"Template_Echo/pkg/config"
	"Template_Echo/pkg/controllers"

	"github.com/labstack/echo/v4"

	"fmt"
)

func main() {
	config.Config()

	appPort := config.AppPort()
	e := echo.New()

	controllers.AppModule(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appPort)))
}
