package internal

import (
	"errors"
	"os"
	"strconv"
)

const (
	DBURLKey         = "DATABASE_URL"
	RedisHostKey     = "REDIS_HOST"
	RedisPasswordKey = "REDIS_PASSWORD"
	RedisPrefixKey   = "REDIS_PREFIX"
	EnvKey           = "ENV"
	PortKey          = "HTTP_PORT"
)

// Config stores a cache of configuration values.
type Config struct {
	cache map[string]interface{}
}

// Error if configuration value is not found
var ErrorConfigNotFound = errors.New("config not found")

// Config attempts to fetch results from it's internal cache,
// if the data is not in the internal cache, it attempts to get it from the env variables,
// if the data is not set in the env variables, it returns a not found error.
func (c *Config) get(name string) (string, error) {
	// Return the value from cache if it is available
	if cacheValue, ok := c.cache[name]; ok {
		if value, ok := cacheValue.(string); ok {
			return value, nil
		}
	}

	// Return the value from the env variables if it is available
	// The value is cached for future queries
	if value, ok := os.LookupEnv(name); ok {
		c.cache[name] = value
		return value, nil
	}

	// Value not available
	return "", ErrorConfigNotFound
}

// Get the postgres database URL
func (c *Config) GetDBURL() (string, error) {
	return c.get(DBURLKey)
}

// Get the redis host
func (c *Config) GetRedisHost() (string, error) {
	return c.get(RedisHostKey)
}

// Get the redis password
func (c *Config) GetRedisPassword() (string, error) {
	return c.get(RedisPasswordKey)
}

// Get the redis prefix
func (c *Config) GetRedisPrefix() (string, error) {
	return c.get(RedisPrefixKey)
}

// Get the env
func (c *Config) GetEnv() (string, error) {
	return c.get(EnvKey)
}

// Get the HTTP Port
func (c *Config) GetHTTPPort() (int, error) {
	val, err := c.get(PortKey)
	if err != nil {
		return 0, err
	}

	port, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return port, nil
}
