package repository

import (
	"errors"
	"fmt"
	"strings"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/repository"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db repository.Database
}

const (
	duplicateEntryMsg = "Duplicate entry"
	numRowInserted    = 1
)

const (
	InsertUserStatement = "INSERT INTO User VALUES('%s','%s','%s','%s')"
)

var (
	insertUserError = errors.New("failed to insert user")
)

func NewUserRepository(db repository.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) Insert(user dto.UserDTO) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(InsertUserStatement,
		user.Username,
		user.Password,
		user.CreateDate,
		user.UpdateDate,
	))

	if err != nil {
		if strings.Contains(err.Error(), duplicateEntryMsg) {
			return repository.DuplicateUser
		}
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if numRowInserted != rowAffected {
		return insertUserError
	}

	return nil
}
