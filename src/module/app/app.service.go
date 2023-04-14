package module_app

import (
	"Template_Echo/src/shared/dto"
	"Template_Echo/src/shared/message"
)

func GetTest() *shared_dto.ResponseDto {

	return &shared_dto.ResponseDto{
		Code: shared_message.SUCCESS,
		Data: nil,
		Message: shared_dto.ResponseMessage{
			Vi: "Thành công",
			En: "Successfully",
		},
	}
}
