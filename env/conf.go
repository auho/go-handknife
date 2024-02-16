package env

import (
	"fmt"

	"github.com/auho/go-handknife/env/form"
	"github.com/auho/go-handknife/env/form/config"
	"github.com/auho/go-handknife/env/form/sh"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ConfigLoad struct {
	former form.Former

	redisMap  map[string]*redis.Client
	mysqlMap  map[string]*gorm.DB
	serverMap map[string]*config.Server
}

func NewConfigLoad() *ConfigLoad {
	return &ConfigLoad{
		redisMap:  make(map[string]*redis.Client),
		mysqlMap:  make(map[string]*gorm.DB),
		serverMap: make(map[string]*config.Server),
	}
}

func (cl *ConfigLoad) FromSh(filePath string) error {
	_sh, err := sh.NewSh(filePath)
	if err != nil {
		return err
	}

	cl.former = _sh

	return nil
}

func (cl *ConfigLoad) LoadEnv() (string, error) {
	return cl.former.LoadEnv()
}

func (cl *ConfigLoad) LoadServer(sName string) (*config.Server, error) {
	if server, ok := cl.serverMap[sName]; ok {
		return server, nil
	}

	server, err := cl.former.LoadServer(sName)
	if err != nil {
		return nil, fmt.Errorf("load server %s; %w", sName, err)
	}

	cl.serverMap[sName] = &server
	return cl.serverMap[sName], nil
}

func (cl *ConfigLoad) LoadMysql(hostName, dbName string) (*gorm.DB, error) {
	key := hostName + dbName
	if _db, ok := cl.mysqlMap[key]; ok {
		return _db, nil
	}

	_db, err := cl.former.LoadMysql(hostName, dbName)
	if err != nil {
		return nil, fmt.Errorf("mysql %s:%s; %w", hostName, dbName, err)
	}

	cl.mysqlMap[key] = _db
	return _db, nil
}

func (cl *ConfigLoad) LoadRedis(hostName string) (*redis.Client, error) {
	if _db, ok := cl.redisMap[hostName]; ok {
		return _db, nil
	}

	_db, err := cl.former.LoadRedis(hostName)
	if err != nil {
		return nil, fmt.Errorf("mysql %s; %w", hostName, err)
	}

	cl.redisMap[hostName] = _db
	return _db, nil
}
