package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Connect to postgres server
func ConnectToPostgres(ctx *AppContext) error {
	config, err := ctx.GetConfig()
	if err != nil {
		return err
	}

	dbUrl, err := config.Get("DATABASE_URL")
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}

	ctx.SetPostgresConn(conn)
	return nil
}
