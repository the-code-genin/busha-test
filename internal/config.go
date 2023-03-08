package internal

import (
	"errors"
	"os"
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
func (c *Config) Get(name string) (string, error) {
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

var DefaultConfig = &Config{
	map[string]interface{}{},
}
