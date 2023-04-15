package controllers

import (
	"Template_Echo/pkg/services"

	"github.com/labstack/echo/v4"
)

func AppController(app *echo.Echo) {
	var appRoute = "/app"
	app.GET(appRoute, services.GetTest)
}
