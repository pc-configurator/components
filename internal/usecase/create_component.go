package usecase

import (
	"context"
	"fmt"

	"github.com/pc-configurator/components/internal/dto"
)

func (u *UseCase) CreateComponent(ctx context.Context, input dto.CreateComponentInput) (dto.CreateComponentOutput, error) {
	var output dto.CreateComponentOutput

	err := input.Validate()
	if err != nil {
		return output, fmt.Errorf("input.Validate: %w", err)
	}

	component, err := u.postgres.CreateComponent(ctx, input)

	output = dto.CreateComponentOutput{
		ID: component.ID,
	}

	return output, nil
}
