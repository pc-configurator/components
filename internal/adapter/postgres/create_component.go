package postgres

import (
	"context"
	"errors"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (p *Postgres) CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error) {
	return domain.Component{}, logger.NewErrorWithPath("penis", errors.New("not implemented"))
}
