package postgres_entities

import (
	"context"

	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
)

func (p *PostgresEntities) CreateCategory(ctx context.Context, input dto.CreateCategoryInput) (domain.Category, error) {
	const sql = `INSERT INTO category (name) VALUES ($1) RETURNING id, name`

	args := []any{
		input.Name,
	}

	var component domain.Category
	err := p.Pool.QueryRow(ctx, sql, args...).Scan(&component.ID, &component.Name)
	if err != nil {
		if postgres.IsConstraint(err, CategoryNameUniqueKey) {
			return component, domain.ErrCategoryNameExists
		}

		return component, logger.NewErrorWithPath("p.Pool.Exec", err)
	}

	return component, nil
}
