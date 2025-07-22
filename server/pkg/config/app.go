package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env variables found")
	}


	log.Println("Connecting to Redis...")
    redisClient= redis.NewClient(&redis.Options{
        Addr:   os.Getenv("REDIS_ADDR") ,
        Password: os.Getenv("REDIS_PASS"), 
        DB:       0,  
    })
	
	log.Println("Connected to Redis")
}

func GetRedis() *redis.Client{
	return redisClient
}