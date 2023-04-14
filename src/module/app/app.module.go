package module_app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppModule(e *echo.Echo) *echo.Echo {

	e.GET("/", func(c echo.Context) error {
		var res = GetTest()
		return c.JSON(http.StatusOK, res)
	})

	return e
}
