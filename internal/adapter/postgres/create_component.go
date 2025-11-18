package postgres

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
)

func (p *Postgres) CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error) {
	return domain.Component{}, nil
}
