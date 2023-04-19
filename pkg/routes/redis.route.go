package routes

import (
	"Template_Echo/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RedisRoute(app *echo.Echo) {
	var redisRoute = "/redis"
	app.GET(redisRoute, controllers.GetRedis)
}
