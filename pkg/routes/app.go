package routes

import (
	"Template_Echo/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func AppRoute(app *echo.Echo) {
	var appRoute = "/app"
	app.GET(appRoute, controllers.GetHello)
}
