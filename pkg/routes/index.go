package routes

import "github.com/labstack/echo/v4"

func Routes(app *echo.Echo) {
	RedisRoute(app)
}
