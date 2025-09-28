package postgres

import "github.com/pc-configurator/components/pkg/postgres"

type Postgres struct {
	pool *postgres.Pool
}

func New(p *postgres.Pool) *Postgres {
	return &Postgres{pool: p}
}
