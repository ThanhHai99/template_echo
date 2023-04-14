package controllers

import (
	"Template_Echo/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppModule(e *echo.Echo) *echo.Echo {

	e.GET("/", func(c echo.Context) error {
		var res = services.GetTest()
		return c.JSON(http.StatusOK, res)
	})

	return e
}
