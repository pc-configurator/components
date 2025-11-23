package domain

import (
	"errors"
)

var (
	ErrValidation          = errors.New("validation_error")
	ErrInternal            = errors.New("internal_error")
	ErrCategoryNotFound    = errors.New("category_not_found")
	ErrComponentNameExists = errors.New("component_name_exists")
	ErrCategoryNameExists  = errors.New("category_name_exists")
	ErrComponentNotFound   = errors.New("component_not_found")
)

type ErrWithDetails struct {
	Code          string             `json:"code"`
	InvalidFields *map[string]string `json:"invalidFields,omitempty"`
}

func NewErrorWithDetails(err error) ErrWithDetails {
	return ErrWithDetails{
		Code: err.Error(),
	}
}

func NewValidationError(invalidFields map[string]string) ErrWithDetails {
	return ErrWithDetails{
		Code:          ErrValidation.Error(),
		InvalidFields: &invalidFields,
	}
}
