package models

import "Template_Echo/pkg/constants"

type ResponseMessage struct {
	Vi string
	En string
}

type ResponseDto struct {
	Code    constants.Code
	Data    interface{}
	Message ResponseMessage
}
