package postgres_entities

import "github.com/pc-configurator/components/pkg/postgres"

type PostgresEntities struct {
	Pool *postgres.Pool
}

func New(p *postgres.Pool) *PostgresEntities {
	return &PostgresEntities{
		Pool: p,
	}
}
