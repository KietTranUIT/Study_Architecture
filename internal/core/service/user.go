package service

import (
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/repository"
	"user-service/internal/core/port/service"
	"user-service/internal/core/common/utils"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (u UserService) SignUp(request *request.SignUpRequest) *response.Response {
	if len(request.Username) == 0 {

		return u.createFailedResponse(error_code.InvalidRequest, error_code.InvalidUsernameMsg)
	}

	if len(request.Password) == 0 {
		return u.createFailedResponse(error_code.InvalidRequest, error_code.InvalidPasswordMsg)
	}

	user := dto.UserDTO {
		Username: request.Username,
		Password: request.Password,
		CreateDate: utils.GetCurrentTime(),
	}

	err := u.userRepo.Insert(user)

	if err != nil {
		if err == repository.DuplicateUser {
			return u.createFailedResponse(error_code.DuplicateUser, error_code.DuplicateUserErrMsg)
		}
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	data := response.SignUpResponse {
		DisplayName: request.Username,
	}
	return u.createSuccessResponse(data)
}

func (u UserService) createFailedResponse(code error_code.ErrorCode, message string) *response.Response {
	return &response.Response{
		Status: false,
		ErrorCode: code,
		ErrorMessage: message,
	}
}

func (u UserService) createSuccessResponse(data response.SignUpResponse) *response.Response {
	return &response.Response{
		Data: data,
		Status: true,
		ErrorCode: error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
	}

}