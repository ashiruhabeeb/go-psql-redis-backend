package db

import (
	"context"
	"log"

	"github.com/ashiruhabeeb/go-backend/pkg/config"

	"github.com/redis/go-redis/v9"
)

func RedisConnect(env *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: env.RedisURI,
		Password: env.RedisPass,
		DB: env.RedisDB,
	})

	if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}

	log.Println("[INIT] ✅ redis client connection established")

	err := rdb.Set(context.TODO(), "test", "Go-psql-redis-backend", 0).Err()
	if err != nil {
		panic(err)
	}

	return rdb
}
