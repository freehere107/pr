package utiles

import (
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
