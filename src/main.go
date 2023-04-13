package main

import (
	"Template_Echo/src/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	appPort := config.AppPort()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World !")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appPort)))
}
