package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (u *UseCase) GetComponent(ctx context.Context, input dto.GetComponentIDInput) (dto.GetComponentIDOutput, error) {
	var output dto.GetComponentIDOutput

	err := input.Validate()
	if err != nil {
		return output, logger.NewErrorWithPath("input.Validate", err)
	}

	outputFromCache, err := u.cacheStorage.GetComponent(ctx, input.ID)
	if err == nil {
		return outputFromCache, nil
	}

	component, err := u.entitiesStorage.GetComponent(ctx, input)
	if err != nil {
		return output, logger.NewErrorWithPath("u.postgres.GetComponent", err)
	}

	output = dto.GetComponentIDOutput{
		Component: component,
	}

	u.cacheStorage.SetComponent(ctx, output)

	return output, nil
}
