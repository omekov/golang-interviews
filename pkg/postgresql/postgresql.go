package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Connection(ctx context.Context, driver, dsn string, timeout time.Duration) (*sqlx.DB, error) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := sqlx.ConnectContext(ctx, driver, dsn)
			if err == nil {
				return db, nil
			}
			log.Println(errors.Wrapf(err, "failed to connect to db %s", dsn))
		}
	}
}
