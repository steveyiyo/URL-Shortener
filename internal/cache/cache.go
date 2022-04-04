package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/steveyiyo/url-shortener/package/tools"
)

var Redis *redis.Client

// Init Client
func InitRedis(Redis_Addr, Redis_Pwd string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     Redis_Addr,
		Password: Redis_Pwd,
		DB:       0,
	})
}

// Add data to Redis
func AddData(key string, value string, second time.Duration) bool {
	ctx := context.Background()

	err := Redis.Set(ctx, key, value, second*time.Second).Err()
	return tools.ErrCheck(err)
}

// Query data from Redis
func QueryData(key string) (bool, string) {
	ctx := context.Background()

	var return_status bool
	var return_value string

	value, err := Redis.Get(ctx, key).Result()
	tools.ErrCheck(err)

	if err == redis.Nil {
		return_status = false
		return_value = ""
	} else if !tools.ErrCheck(err) {
		log.Println(err)
	} else {
		return_status = true
		return_value = value
	}
	return return_status, return_value
}
