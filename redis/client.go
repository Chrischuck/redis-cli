package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func createRedisClient() *redis.Client {

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       redisDB,                     // use default DB
	})

	return client
}

// Client is the redis client
var Client = createRedisClient()
