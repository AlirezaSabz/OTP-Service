package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func Connect(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	return rdb
}
