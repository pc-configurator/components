package usecase

import (
	"context"

	"github.com/pc-configurator/components/internal/entity"
)

type Postgres interface {
	SelectComponent(ctx context.Context, componentID int) (entity.Component, error)
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
