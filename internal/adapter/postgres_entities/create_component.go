package postgres_entities

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
)

func (p *PostgresEntities) CreateComponent(ctx context.Context, input dto.CreateComponentInput) (domain.Component, error) {
	const sql = `INSERT INTO component (name, price, category_id, description) VALUES ($1, $2, $3, $4) RETURNING id, name, price, category_id, description`

	args := []any{
		input.Name,
		input.Price,
		input.CategoryID,
		input.Description,
	}

	var component domain.Component
	err := p.Pool.QueryRow(ctx, sql, args...).Scan(&component.ID, &component.Name, &component.Price, &component.CategoryID, &component.Description)
	if err != nil {
		if postgres.IsConstraint(err, ComponentNameUniqueKey) {
			return domain.Component{}, domain.ErrComponentNameExists
		}

		if postgres.IsConstraint(err, ComponentCategoryForeignKey) {
			return domain.Component{}, domain.ErrCategoryNotFound
		}

		return component, logger.NewErrorWithPath("p.Pool.Exec", err)
	}

	return component, nil
}
