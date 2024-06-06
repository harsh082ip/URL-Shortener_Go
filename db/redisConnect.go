package db

import "github.com/go-redis/redis"

func RedisConnect() *redis.Client {
	url := "redis://localhost:6379"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
