package usecase

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
	ErrInternal   = errors.New("internal error")
)
