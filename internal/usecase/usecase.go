package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
)

type Postgres interface {
	CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error)
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{postgres: p}
}
