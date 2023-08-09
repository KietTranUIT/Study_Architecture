package service

import (
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

type UserService interface {
	SignUp(req *request.SignUpRequest) *response.Response
}
