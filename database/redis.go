package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

type RedisInstance struct {
	Client *redis.Client
}

var Redis RedisInstance

func ConnectRedis() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DNS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	if err != nil {
		log.Fatal("Failed to connect to redis. \n", err)
	}

	Redis = RedisInstance{
		Client: client,
	}

}
