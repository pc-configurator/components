package dto

import (
	"github.com/pc-configurator/components/pkg/validation"
)

type CreateComponentInput struct {
	Name        *string `json:"name"`
	Price       *int    `json:"price"`
	Description *string `json:"description"`
	CategoryID  *int    `json:"categoryId"`
}

func (input CreateComponentInput) Validate() error {
	errors := validation.ErrorFields{}

	if input.Name == nil {
		errors["name"] = "name is required"
	}

	if input.Price == nil {
		errors["price"] = "price is required"
	}

	if input.Description == nil {
		errors["description"] = "description is required"
	}

	if input.CategoryID == nil {
		errors["categoryId"] = "categoryId is required"
	}

	if input.Name != nil && !validation.MinString(input.Name, 4) {
		errors["name"] = "name must be at least 4 characters long"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type CreateComponentOutput struct {
	ID int
}
