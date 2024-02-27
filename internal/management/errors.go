package management

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrInternalError = errors.New("internal error")
)

func GormToCustomError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrAlreadyExists
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound
	default:
		return ErrInternalError
	}
}

func CustomErrorToHTTPStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch {
	case errors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
