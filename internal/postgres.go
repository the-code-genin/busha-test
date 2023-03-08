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

	dbUrl, err := config.GetDBURL()
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}

	ctx.setPostgresConn(conn)
	return nil
}
