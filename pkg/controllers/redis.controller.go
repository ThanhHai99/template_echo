package controllers

import (
	"Template_Echo/pkg/constants"
	"Template_Echo/pkg/models"
	"Template_Echo/pkg/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRedis(c echo.Context) error {
	data := services.GetHello()
	res := &models.ResponseDto{
		Code: constants.SUCCESS,
		Data: data,
		Message: models.ResponseMessage{
			Vi: "Thành công",
			En: "Successfully",
		},
	}

	return c.JSON(http.StatusOK, res)
}

func SetRedis(c echo.Context) error {
	fmt.Println(c.Request().Header)
	fmt.Println(c.QueryParams())
	fmt.Println(c.Request().Body)
	data := services.GetHello()
	res := &models.ResponseDto{
		Code: constants.SUCCESS,
		Data: data,
		Message: models.ResponseMessage{
			Vi: "Thành công",
			En: "Successfully",
		},
	}

	return c.JSON(http.StatusOK, res)
}
