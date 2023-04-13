package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AppModule(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET hehe boi")
	})
	return e
}
