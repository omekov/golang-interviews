package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func Connection(ctx context.Context, driver, uri string) (*sqlx.DB, error) {
	return sqlx.ConnectContext(ctx, driver, uri)
}
