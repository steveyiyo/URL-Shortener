package cache_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	Tools "github.com/steveyiyo/url-shortener/internal/tools"
)

var Redis *redis.Client

// Init Client
func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

// Add data to Redis
func AddData(key, value string) bool {
	ctx := context.Background()

	err := Redis.Set(ctx, key, value, 30*time.Second).Err()
	return Tools.ErrCheck(err)
}

// Query data from Redis
func QueryData(key string) (bool, string) {
	ctx := context.Background()

	var return_status bool
	var return_value string

	val2, err := Redis.Get(ctx, key).Result()
	Tools.ErrCheck(err)

	if err == redis.Nil {
		return_status = false
		return_value = ""
	} else if !Tools.ErrCheck(err) {
		log.Println(err)
	} else {
		return_status = true
		return_value = val2
	}
	return return_status, return_value
}

// It's a test function.
func TestMain(t *testing.T) {
	InitRedis()
	AddData("hi", "pong")
	status, data := QueryData("hi")
	if status {
		fmt.Println(data)
	} else {
		fmt.Println("QaQ")
	}
}
