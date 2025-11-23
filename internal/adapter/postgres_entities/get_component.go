package postgres_entities

import (
	"context"
	"errors"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
)

func (p *PostgresEntities) GetComponent(ctx context.Context, input dto.GetComponentIDInput) (domain.Component, error) {
	const sql = `SELECT id, name, price, category_id, description FROM component WHERE id = $1`

	var component domain.Component
	err := p.Pool.QueryRow(ctx, sql, input.ID).Scan(&component.ID, &component.Name, &component.Price, &component.CategoryID, &component.Description)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.Component{}, domain.ErrComponentNotFound
		}

		return domain.Component{}, logger.NewErrorWithPath("p.Pool.Exec", err)
	}

	return component, nil
}
