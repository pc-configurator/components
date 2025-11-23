package dto

import "github.com/pc-configurator/components/pkg/validation"

type CreateCategoryInput struct {
	Name *string
}

func (input CreateCategoryInput) Validate() error {
	errors := validation.ErrorFields{}

	if input.Name == nil {
		errors["name"] = "name is required"
	}

	if input.Name != nil && !validation.MinString(input.Name, 4) {
		errors["name"] = "name must be at least 4 characters long"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type CreateCategoryOutput struct {
	ID int
}
