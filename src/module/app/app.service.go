package module_app

import (
	shared_dto "Template_Echo/src/shared/dto"
	shared_message "Template_Echo/src/shared/message"
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
