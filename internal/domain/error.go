package domain

import (
	"errors"
)

var (
	ErrValidation       = errors.New("validation_error")
	ErrInternal         = errors.New("internal_error")
	ErrCategoryNotFound = errors.New("category_not_found")
)

type ErrWithDetails struct {
	Code          string             `json:"code"`
	InvalidFields *map[string]string `json:"invalidFields,omitempty"`
}

func NewValidationError(invalidFields map[string]string) ErrWithDetails {
	return ErrWithDetails{
		Code:          ErrValidation.Error(),
		InvalidFields: &invalidFields,
	}
}

func NewInternalError() ErrWithDetails {
	return ErrWithDetails{
		Code: ErrInternal.Error(),
	}
}

func NewCategoryNotFoundError() ErrWithDetails {
	return ErrWithDetails{
		Code: ErrCategoryNotFound.Error(),
	}
}
