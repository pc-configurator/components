package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
)

type CacheStorage interface {
	GetComponent(ctx context.Context, componentID string) (dto.GetComponentIDOutput, error)
	SetComponent(ctx context.Context, component dto.GetComponentIDOutput)
}

type EntitiesStorage interface {
	CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error)
	CreateCategory(ctx context.Context, input dto.CreateCategoryInput) (domain.Category, error)
	GetComponent(ctx context.Context, input dto.GetComponentIDInput) (domain.Component, error)
}

type UseCase struct {
	entitiesStorage EntitiesStorage
	cacheStorage    CacheStorage
}

func New(e EntitiesStorage, c CacheStorage) *UseCase {
	return &UseCase{entitiesStorage: e, cacheStorage: c}
}
