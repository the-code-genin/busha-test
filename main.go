package main

import (
	"context"

	"github.com/inconshreveable/log15"
	"github.com/the-code-genin/busha-test/internal"
)

// Context encapsulates redis and postgres connections
var ctx context.Context

// Setup application context
func init() {
	// Setup initial context
	ctx = context.TODO()

	// Load env variables into memory
	log15.Info("Loading env variables")
	if err := internal.LoadEnvVariables(); err != nil {
		panic(err)
	}

	// Connect to postgres db
	log15.Info("Connecting to postgres database")
	pgConn, err := internal.ConnectToPostgres()
	if err != nil {
		panic(err)
	}
	ctx = internal.SetPostgresConn(ctx, pgConn)

	// Connect to redis
	log15.Info("Connecting to redis server")
	redisClient, err := internal.ConnectToRedis()
	if err != nil {
		panic(err)
	}
	ctx = internal.SetRedisClient(ctx, redisClient)

	// Seed the application database at first startup
	log15.Info("Seeding system with swapi data")
	if err = SeedDatabase(ctx); err != nil {
		panic(err)
	}
}

func main() {
	// Main
}
