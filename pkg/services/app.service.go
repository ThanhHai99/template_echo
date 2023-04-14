package services

import (
	"Template_Echo/pkg/constants"
	"Template_Echo/pkg/models"
)

func GetTest() *models.ResponseDto {

	return &models.ResponseDto{
		Code: constants.SUCCESS,
		Data: nil,
		Message: models.ResponseMessage{
			Vi: "Thành công",
			En: "Successfully",
		},
	}
}
