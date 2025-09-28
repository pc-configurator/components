package dto

import "github.com/pc-configurator/components/internal/entity"

type GetComponentInput struct {
	ID string `json:"id" validate:"required,numeric"`
}

type GetComponentOutput struct {
	entity.Component
}
