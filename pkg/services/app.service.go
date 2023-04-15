package services

import (
	"Template_Echo/pkg/constants"
	"Template_Echo/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTest(c echo.Context) error {
	res := &models.ResponseDto{
		Code: constants.SUCCESS,
		Data: nil,
		Message: models.ResponseMessage{
			Vi: "Thành công",
			En: "Successfully",
		},
	}

	return c.JSON(http.StatusOK, res)
}
