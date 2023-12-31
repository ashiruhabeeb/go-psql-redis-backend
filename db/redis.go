package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient	*redis.Client
)

func RedisConnect(addr, pwd string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: pwd,
		DB: db,
	})

	if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
		log.Fatalf("[ERROR] rdb.Ping failure: %v", err)
	}

	log.Println("[INIT] ✅ redis client connection established")

	err := rdb.Set(context.TODO(), "test", "Go-psql-redis-backend", 0).Err()
	if err != nil {
		panic(err)
	}

	return rdb
}
