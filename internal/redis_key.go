package internal

import "fmt"

// Prefixes the key with the app redis key for namespacing
func RedisKey(key string) string {
	redisPrefix, err := DefaultConfig.Get("REDIS_PREFIX")
	if err != nil {
		return key
	}
	return fmt.Sprintf("%s.%s", redisPrefix, key)
}
