package usecase

import (
	"context"
	"strconv"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/internal/entity"
	"github.com/pc-configurator/components/pkg/base_errors"
)

func (uc *UseCase) GetComponent(ctx context.Context, input dto.GetComponentInput) (dto.GetComponentOutput, error) {
	ID, err := strconv.Atoi(input.ID)
	if err != nil {
		return dto.GetComponentOutput{}, entity.ErrInvalidID
	}

	component, err := uc.postgres.SelectComponent(ctx, ID)
	if err != nil {
		return dto.GetComponentOutput{}, base_errors.WithPath("postgres.SelectComponent", err)
	}

	return dto.GetComponentOutput{Component: component}, nil
}
