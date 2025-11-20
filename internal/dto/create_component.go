package dto

import (
	"github.com/pc-configurator/components/pkg/validation"
)

type CreateComponentInput struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (input CreateComponentInput) Validate() error {
	errors := validation.ErrorFields{}

	if !validation.MinString(input.Name, 4) {
		errors["name"] = "Name must be at least 4 characters long"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type CreateComponentOutput struct {
	ID int
}
