package config

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONN"),
		Password: "",
		DB:       0,
	})
	return client
}

func DeleteRedisData(token string) error {
	rdb := NewRedisClient()
	err := rdb.Del(context.Background(), token).Err()
	if err != nil {
		return errors.New("failed to delete token in redis")
	}

	return nil
}

func SetRedisData(token, userUuid string) error {
	rdb := NewRedisClient()
	err := rdb.Set(context.Background(), token, userUuid, 24*time.Hour).Err()
	if err != nil {
		return errors.New("failed to save token in Redis")

	}

	return nil
}

func GetRedisData(splitToken []string) error {
	rdb := NewRedisClient()
	err := rdb.Get(context.Background(), splitToken[1]).Err()
	if err != nil {
		return errors.New("unauthorized")
	}
	return nil
}
