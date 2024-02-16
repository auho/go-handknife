package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Application struct {
	redisHub map[string]*redis.Client
	dbHub    map[string]*gorm.DB
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

func (a *Application) GetRedis(conf RedisConf) *redis.Client {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if r, ok := a.redisHub[conf.Name]; !ok {
		var err error
		r, err = a.openRedis(conf)
		if err != nil {
			log.Fatal(err)
		}

		a.redisHub[conf.Name] = r

		return r
	} else {
		return r
	}
}

func (a *Application) GetDB(conf MysqlConf) *gorm.DB {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if d, ok := a.dbHub[conf.Name]; !ok {
		var err error
		d, err = a.openMysql(conf)
		if err != nil {
			log.Fatal(err)
		}

		a.dbHub[conf.Name] = d

		return d
	} else {
		return d
	}
}

func (a *Application) GetEs(conf EsConf) *elasticsearch.Client {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if e, ok := a.esHub[conf.Name]; !ok {
		var err error
		e, err = a.openEs(conf)
		if err != nil {
			log.Fatal(err)
		}

		a.esHub[conf.Name] = e

		return e
	} else {
		return e
	}
}

func (a *Application) openRedis(conf RedisConf) (*redis.Client, error) {
	if conf.Host == "" {
		return nil, errors.New(fmt.Sprintf("redis[%s] addr is empty", conf.Name))
	}

	if conf.Port <= 0 {
		conf.Port = 6379
	}

	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Auth,
		DB:       conf.DB,
		PoolSize: 1,
	})

	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis[%s] ping error %w", conf.Name, err)
	}

	return redisClient, nil
}

func (a *Application) openMysql(conf MysqlConf) (*gorm.DB, error) {
	if conf.Host == "" {
		return nil, errors.New(fmt.Sprintf("msyql[%s] host is empty", conf.Name))
	}

	if conf.Port <= 0 {
		conf.Port = 3306
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
			logger.Config{
				SlowThreshold: time.Second,  // 慢 SQL 阈值
				LogLevel:      logger.Error, // 日志级别
			},
		)})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("msyql[%s] %v", conf.Name, err))
	}

	err = db.Select("select version()").Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("msyql[%s] connection %v", conf.Name, err))
	}

	return db, nil
}

func (a *Application) openEs(conf EsConf) (*elasticsearch.Client, error) {
	if len(conf.Address) <= 0 {
		return nil, errors.New(fmt.Sprintf("es[%s] address is empty", conf.Name))
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:  conf.Address,
		Username:   conf.Username,
		Password:   conf.Password,
		MaxRetries: 3,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("elasticsearch[%s] client %v", conf.Name, err))
	}

	resp, err := es.Ping()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("elasticsearch[%s] ping error %v", conf.Name, err))
	}

	if resp.IsError() {
		return nil, errors.New(fmt.Sprintf("elasticsearch[%s] ping error %s", conf.Name, resp.Status()))
	}

	return es, nil
}
