package utils

import (
	"rumah-sakit/model/dto"
)

func HandleResponse(data interface{}, message string) dto.Response {
	response := dto.Response{
		Data:    data,
		Message: message,
	}

	return response
}
