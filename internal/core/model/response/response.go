package response

import (
	"user-service/internal/core/entity/error_code"
)

type Response struct {
	Data         interface{}          `json:"data"`
	Status       bool                 `json:"status"`
	ErrorCode    error_code.ErrorCode `json:"error_code"`
	ErrorMessage string               `json:"errorMessage"`
}

type SignUpResponse struct {
	DisplayName string
}
