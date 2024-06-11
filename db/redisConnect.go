package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	REDIS_URI := os.Getenv("REDIS_URI")

	opts, err := redis.ParseURL(REDIS_URI)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
