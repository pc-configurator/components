package dto

import "errors"

type CreateComponentInput struct {
	Name        string
	Price       int
	Category    string
	Description string
}

func (input CreateComponentInput) Validate() error {
	if len([]rune(input.Name)) < 4 {
		return errors.New("длина должна быть больше 3")
	}

	return nil
}

type CreateComponentOutput struct {
	ID int
}
