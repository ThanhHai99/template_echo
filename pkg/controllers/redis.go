package controllers

import (
	"Template_Echo/pkg/constants"
	"Template_Echo/pkg/models"
	"Template_Echo/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetData(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	data := services.GetHello(page, limit)
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
