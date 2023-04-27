package routes

import (
	"Template_Echo/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(app *echo.Echo) {
	var appRoute = "/user"
	app.GET(appRoute, controllers.ReadAllUser)
	app.GET(appRoute+"/:id", controllers.ReadOneUser)
}
