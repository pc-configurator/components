package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsConstraint(err error, constraintKey string) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr.ConstraintName == constraintKey
	}

	return false
}
