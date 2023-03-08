package main

import (
	"context"
	"sync"

	"github.com/inconshreveable/log15"
	"github.com/the-code-genin/busha-test/api"
	"github.com/the-code-genin/busha-test/internal"
)

// Context encapsulating application resources
var ctx *internal.AppContext
var wg sync.WaitGroup

// Setup application context
func init() {
	// Setup initial context
	ctx := internal.NewAppContext(context.TODO())

	// Load env variables into memory
	log15.Info("Loading env variables")
	if err := internal.LoadEnvVariables(ctx); err != nil {
		panic(err)
	}

	// Connect to postgres db
	log15.Info("Connecting to postgres database")
	if err := internal.ConnectToPostgres(ctx); err != nil {
		panic(err)
	}

	// Connect to redis
	log15.Info("Connecting to redis server")
	if err := internal.ConnectToRedis(ctx); err != nil {
		panic(err)
	}

	// Seed the application database at first startup
	log15.Info("Seeding system with swapi data")
	if err := SeedDatabase(ctx); err != nil {
		panic(err)
	}
}

func main() {
	// Setup API server
	server, err := api.NewServer(ctx)
	if err != nil {
		panic(err)
	}

	// Run API server
	wg.Add(1)
	go func() {
		log15.Info("Serving HTTP requests")
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
