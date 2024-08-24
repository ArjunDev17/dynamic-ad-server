package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitConfig() {
	// Determine the environment and load the corresponding .env file
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddr := redisHost + ":" + redisPort

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"), // No password set
	})

	// Test the Redis connection
	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
