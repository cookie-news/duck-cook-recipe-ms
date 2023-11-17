package redis

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func Connect() *redis.Client {
	hostStr := os.Getenv("REDIS_HOST")
	passStr := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")

	if dbStr == "" {
		log.Println("Warning: REDIS_DB environment variable is not set. Using default value 0.")
		dbStr = "0"
	}

	dbInt, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatalf("Failed to convert REDIS_DB to integer: %v", err)
	}

	log.Println("Redis Host: " + hostStr)

	client := redis.NewClient(&redis.Options{
		Addr:     hostStr,
		Password: passStr,
		DB:       dbInt,
	})

	client = client.WithTimeout(2 * time.Second)

	return client
}
