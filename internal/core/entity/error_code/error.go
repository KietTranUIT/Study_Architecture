package error_code

type ErrorCode string

const (
	Success        ErrorCode = "SUCCESS"
	InvalidRequest ErrorCode = "INVALID_REQUEST"
	DuplicateUser  ErrorCode = "DUPLICATE_USER"
	InternalError  ErrorCode = "INTERNAL_ERROR"
)

type ErrorMsg string

const (
	SuccessErrMsg       = "success"
	InvalidUsernameMsg  = "invalid username"
	InvalidPasswordMsg  = "invalid password"
	DuplicateUserErrMsg = "duplicate user"
	InternalErrMsg      = "internal error"
)
