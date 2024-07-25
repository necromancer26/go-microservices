package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	rb *redis.Client
}

var ctx = context.Background()

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use your Redis server address and port
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Ping the Redis server to ensure connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis:")
	}

	return rdb
}

// GetAllKeysAndValues retrieves all keys and their values from Redis
func GetAllKeysAndValues(rdb *redis.Client) (map[string]string, error) {
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		values[key] = val
	}

	return values, nil
}
