package shared_dto

import "Template_Echo/src/shared/message"

type ResponseMessage struct {
	Vi string
	En string
}

type ResponseDto struct {
	Code    shared_message.Code
	Data    interface{}
	Message ResponseMessage
}
