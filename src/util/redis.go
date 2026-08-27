package util

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx context.Context

func OpenRedis() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT")),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       0,
	})
}

func GetRedisContext() context.Context {
	return ctx
}

func GetDefaultRedisExpiration() time.Duration {
	return time.Duration(time.Duration.Seconds(60))
}

func GetRedis() *redis.Client {
	return rdb
}

func CloseRedis() {
	rdb.Close()
}
