package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pc-configurator/components/internal/entity"
	"github.com/pc-configurator/components/pkg/base_errors"
	"github.com/pc-configurator/components/pkg/postgres"
)

func (p *Postgres) SelectComponent(ctx context.Context, componentID int) (entity.Component, error) {
	const sql = `SELECT id, name FROM component WHERE id = $1`

	dto := struct {
		ID   pgtype.Int4
		Name pgtype.Text
	}{}

	dest := []any{
		&dto.ID,
		&dto.Name,
	}

	err := p.pool.QueryRow(ctx, sql, componentID).Scan(dest...)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return entity.Component{}, base_errors.NotFound
		}

		return entity.Component{}, base_errors.WithPath("pool.QueryRow", err)
	}

	return entity.Component{ID: int(dto.ID.Int32), Name: dto.Name.String}, nil
}
