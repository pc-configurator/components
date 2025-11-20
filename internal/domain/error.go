package domain

import (
	"errors"
)

var (
	ValidationError = errors.New("validation_error")
	InternalError   = errors.New("internal_error")
)

type ErrorWithDetails struct {
	Code          string             `json:"code"`
	InvalidFields *map[string]string `json:"invalidFields,omitempty"`
}

func NewValidationError(invalidFields map[string]string) ErrorWithDetails {
	return ErrorWithDetails{
		Code:          ValidationError.Error(),
		InvalidFields: &invalidFields,
	}
}

func NewInternalError() ErrorWithDetails {
	return ErrorWithDetails{
		Code: InternalError.Error(),
	}
}
