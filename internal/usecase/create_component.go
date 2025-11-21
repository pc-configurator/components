package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (u *UseCase) CreateComponent(ctx context.Context, input dto.CreateComponentInput) (dto.CreateComponentOutput, error) {
	var output dto.CreateComponentOutput

	err := input.Validate()
	if err != nil {
		return output, logger.NewErrorWithPath("input.Validate", err)
	}

	component, err := u.entitiesStorage.CreateComponent(ctx, input)
	if err != nil {
		return output, logger.NewErrorWithPath("u.postgres.CreateComponent", err)
	}

	output = dto.CreateComponentOutput{
		ID: component.ID,
	}

	return output, nil
}
