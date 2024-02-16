package form

import (
	"github.com/auho/go-handknife/env/form/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Former interface {
	LoadEnv() (string, error)
	LoadServer(name string) (config.Server, error)
	LoadMysql(name, dbName string) (*gorm.DB, error)
	LoadRedis(name string) (*redis.Client, error)
}
