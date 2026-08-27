package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func OpenRedis() {
	if rdb != nil {
		return
	}
	cacheHost := os.Getenv("CACHE_HOST")
	if cacheHost == "" {
		cacheHost = "127.0.0.1"
	}
	cachePort := os.Getenv("CACHE_PORT")
	if cachePort == "" {
		cachePort = "6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheHost, cachePort),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("[Redis Warning] Could not connect to Redis at %s:%s: %v", cacheHost, cachePort, err)
	} else {
		log.Printf("[Redis] Connected successfully: %s", pong)
	}
}

func GetRedisContext() context.Context {
	return ctx
}

func GetDefaultRedisExpiration() time.Duration {
	return 60 * time.Second
}

func GetRedis() *redis.Client {
	return rdb
}

func CloseRedis() {
	if rdb != nil {
		_ = rdb.Close()
		rdb = nil
	}
}
