package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Addr     string
	Port     string
	Username string
	Password string
}

func LoadRedis(name string, rc RedisConf) (*redis.Client, error) {
	if rc.Addr == "" {
		return nil, errors.New(fmt.Sprintf("redis[%s] addr is empty", name))
	}

	if rc.Port == "" {
		rc.Port = "6379"
	}

	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     rc.Addr + ":" + rc.Port,
		Username: rc.Username,
		Password: rc.Password,
		DB:       0,
		PoolSize: 1,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis[%s] ping error %w", name, err)
	}

	return redisClient, nil
}
