package config

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

func NewRedisClient() *RedisClient {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		password = ""
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Println("Invalid REDIS_DB, defaulting to 0")
		db = 0
	}

	poolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	if err != nil {
		log.Println("Invalid REDIS_POOL_SIZE, defaulting to 10")
		poolSize = 10
	}

	return &RedisClient{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	}
}

func (r *RedisClient) Connect() *redis.Client {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return client
}
