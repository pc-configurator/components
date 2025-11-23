package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
)

type EntitiesStorage interface {
	CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error)
	CreateCategory(ctx context.Context, input dto.CreateCategoryInput) (domain.Category, error)
}

type UseCase struct {
	entitiesStorage EntitiesStorage
}

func New(e EntitiesStorage) *UseCase {
	return &UseCase{entitiesStorage: e}
}
