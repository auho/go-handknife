package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConf struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (a *Application) GetDB(ctx context.Context, conf MysqlConf) (*gorm.DB, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if d, ok := a.dbHub[conf.Name]; ok {
		return d, nil
	}

	d, err := a.openMysql(ctx, conf)
	if err != nil {
		return nil, err
	}

	a.dbHub[conf.Name] = d

	return d, nil
}

func (a *Application) openMysql(_ context.Context, conf MysqlConf) (*gorm.DB, error) {
	if conf.Host == "" {
		return nil, fmt.Errorf("mysql[%s] host is empty", conf.Name)
	}

	if conf.Port <= 0 {
		conf.Port = 3306
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Error,
			},
		)})
	if err != nil {
		return nil, fmt.Errorf("mysql[%s] %w", conf.Name, err)
	}

	err = db.Select("select version()").Error
	if err != nil {
		return nil, fmt.Errorf("mysql[%s] connection %w", conf.Name, err)
	}

	return db, nil
}
