package main

import (
	"Template_Echo/pkg/configs"
	"Template_Echo/pkg/controllers"
	"Template_Echo/pkg/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	controllers.Controllers(app)
	configs.Configs()
	appPort := configs.AppPort()

	logConfig := utils.DefaultLoggerConfig
	logConfig.Output = utils.Log()
	logConfig.Format = `${remote_ip} ${data_in_out} | ${method}:${uri} | ${status} | ${latency_human} | ${error}`

	app.Use(utils.LoggerWithConfig(logConfig))
	app.Use(middleware.Recover())

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appPort)))
}
