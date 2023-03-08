package internal

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type ContextKey int

const (
	redisContextKey    ContextKey = 1
	postgresContextKey ContextKey = 2
)

// Set redis client on the context
func SetRedisClient(ctx context.Context, redisClient *redis.Client) context.Context {
	return context.WithValue(ctx, redisContextKey, redisClient)
}

// Extract redis client from the context
func GetRedisClient(ctx context.Context) (*redis.Client, error) {
	redisClient, ok := ctx.Value(redisContextKey).(*redis.Client)
	if !ok {
		return nil, errors.New("redis client not found in context")
	}
	return redisClient, nil
}

// Set postgres connection on the context
func SetPostgresConn(ctx context.Context, pgConn *pgx.Conn) context.Context {
	return context.WithValue(ctx, postgresContextKey, pgConn)
}

// Extract postgres connection from the context
func GetPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	pgConn, ok := ctx.Value(postgresContextKey).(*pgx.Conn)
	if !ok {
		return nil, errors.New("postgres connection not found in context")
	}
	return pgConn, nil
}
