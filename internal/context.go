package internal

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type ContextKey int

const (
	configContextKey   ContextKey = 1
	redisContextKey    ContextKey = 2
	postgresContextKey ContextKey = 3
)

type AppContext struct {
	ctx context.Context
}

// Set configuration on the context
func (s *AppContext) setConfig(config *Config) {
	s.ctx = context.WithValue(s.ctx, configContextKey, config)
}

// Extract configuration from the context
func (s *AppContext) GetConfig() (*Config, error) {
	config, ok := s.ctx.Value(configContextKey).(*Config)
	if !ok {
		return nil, errors.New("configuration not found in context")
	}
	return config, nil
}

// Set redis client on the context
func (s *AppContext) setRedisClient(redisClient *redis.Client) {
	s.ctx = context.WithValue(s.ctx, redisContextKey, redisClient)
}

// Extract redis client from the context
func (s *AppContext) GetRedisClient() (*redis.Client, error) {
	redisClient, ok := s.ctx.Value(redisContextKey).(*redis.Client)
	if !ok {
		return nil, errors.New("redis client not found in context")
	}
	return redisClient, nil
}

// Set postgres connection on the context
func (s *AppContext) setPostgresConn(pgConn *pgx.Conn) {
	s.ctx = context.WithValue(s.ctx, postgresContextKey, pgConn)
}

// Extract postgres connection from the context
func (s *AppContext) GetPostgresConn() (*pgx.Conn, error) {
	pgConn, ok := s.ctx.Value(postgresContextKey).(*pgx.Conn)
	if !ok {
		return nil, errors.New("postgres connection not found in context")
	}
	return pgConn, nil
}

func NewAppContext(ctx context.Context) *AppContext {
	return &AppContext{ctx}
}
