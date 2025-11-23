package dto

import (
	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/pkg/validation"
)

type GetComponentIDInput struct {
	ID string
}

func (input GetComponentIDInput) Validate() error {
	errors := validation.ErrorFields{}

	if !validation.IsNumericString(input.ID) {
		errors["id"] = "ID should be a numeric string"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type GetComponentIDOutput struct {
	domain.Component
}
