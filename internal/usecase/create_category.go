package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (u *UseCase) CreateCategory(ctx context.Context, input dto.CreateCategoryInput) (dto.CreateCategoryOutput, error) {
	var output dto.CreateCategoryOutput

	err := input.Validate()
	if err != nil {
		return output, logger.NewErrorWithPath("input.Validate", err)
	}

	category, err := u.entitiesStorage.CreateCategory(ctx, input)
	if err != nil {
		return output, logger.NewErrorWithPath("u.postgres.CreateCategory", err)
	}

	output = dto.CreateCategoryOutput{
		ID: category.ID,
	}

	return output, nil
}
