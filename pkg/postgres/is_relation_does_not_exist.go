package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsRelationDoesNotExist(err error, relationKey string) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.ConstraintName == relationKey
	}

	return false
}
