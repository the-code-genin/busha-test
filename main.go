package main

import (
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/the-code-genin/busha-test/internal"
)

var pgConn *pgx.Conn
var redisClient *redis.Client

func init() {
	// Load env variables into memory
	err := internal.LoadEnvVariables()
	if err != nil {
		panic(err)
	}

	// Connect to postgres db
	pgConn, err = internal.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	// Connect to redis
	redisClient, err = internal.ConnectToRedis()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Main
}