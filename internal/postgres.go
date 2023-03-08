package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Connect to postgres server
func ConnectToPostgres() (*pgx.Conn, error) {
	dbUrl, err := DefaultConfig.Get("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
