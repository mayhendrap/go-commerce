package helpers

import (
	"go-commerce/dtos/response"
	"net/http"
)

var ResponseInternalServerError = ToDefaultResponse(http.StatusInternalServerError, "Failed", "internal server error")

func ToDefaultResponse(status int, message string, data interface{}) response.DefaultResponse {
	return response.DefaultResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
