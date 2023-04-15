package main

import (
	"Template_Echo/pkg/config"
	"Template_Echo/pkg/controllers"
	"Template_Echo/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fmt"
)

func main() {
	config.Config()
	appPort := config.AppPort()

	app := echo.New()
	logConfig := utils.DefaultLoggerConfig
	logConfig.Output = utils.Log()
	logConfig.Format = `${remote_ip} ${data_in_out} | ${method}:${uri} | ${status} | ${latency_human} | ${error}`

	app.Use(utils.LoggerWithConfig(logConfig))

	app.Use(middleware.Recover())

	controllers.AppController(app)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appPort)))
}
