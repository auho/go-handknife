package app

import (
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Application struct {
	dbHub    map[string]*gorm.DB
	redisHub map[string]*redis.Client
	esHub    map[string]*elasticsearch.Client

	mutex sync.Mutex
}

func NewApplication() *Application {
	a := &Application{}
	a.redisHub = make(map[string]*redis.Client)
	a.dbHub = make(map[string]*gorm.DB)
	a.esHub = make(map[string]*elasticsearch.Client)

	return a
}
