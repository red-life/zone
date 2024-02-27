package management

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrInternalError = errors.New("internal error")
)

func GormToCustomError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrAlreadyExists
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound
	default:
		return ErrInternalError
	}
}
