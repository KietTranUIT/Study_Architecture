package repository

import (
	"errors"
	"user-service/internal/core/dto"
)

var (
	DuplicateUser = errors.New("Duplicate user")
)

type UserRepository interface {
	Insert(user dto.UserDTO) error
}
