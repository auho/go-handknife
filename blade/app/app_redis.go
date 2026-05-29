package app

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Auth string `yaml:"auth"`
	DB   int    `yaml:"db"`

	poolSize int
}

func (a *Application) GetRedis(ctx context.Context, conf RedisConf) (*redis.Client, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if r, ok := a.redisHub[conf.Name]; ok {
		return r, nil
	}

	r, err := a.openRedis(ctx, conf)
	if err != nil {
		return nil, err
	}

	a.redisHub[conf.Name] = r

	return r, nil
}

func (a *Application) openRedis(ctx context.Context, conf RedisConf) (*redis.Client, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("redis[%s] addr is empty", conf.Name)
	}

	if conf.Port <= 0 {
		conf.Port = 6379
	}

	poolSize := conf.poolSize
	if poolSize <= 0 {
		poolSize = 1
	}

	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Auth,
		DB:       conf.DB,
		PoolSize: poolSize,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis[%s] ping error %w", conf.Name, err)
	}

	return redisClient, nil
}
