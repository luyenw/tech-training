package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
)

var (
	instance   *redis.Client
	err        error
	once       sync.Once
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PWD  = os.Getenv("REDIS_PWD")
)

func getRedis() *redis.Client {
	fmt.Println(REDIS_HOST)
	return redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: REDIS_PWD,
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	if instance == nil {
		once.Do(func() {
			instance = getRedis()
		})
	}
	return instance
}
