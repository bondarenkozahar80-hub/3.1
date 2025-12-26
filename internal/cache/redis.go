package cache

import (
	"context"
	"delayed-notifier/internal/config"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client redis.Client
}

func New() *Redis {
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Host + config.Cfg.Redis.Port,
		Password: password,
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	fmt.Println("Connected to Redis:", pong)

	return &Redis{
		client: *client,
	}
}

func (r *Redis) Set(key int, value interface{}) error {
	id := strconv.Itoa(key)

	return r.client.SetEx(context.Background(), id, value, 24*time.Hour).Err()
}

func (r *Redis) Get(key string) (string, error) {
	value, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		log.Fatal("error get key: %w", err)
	}

	return value, nil
}
