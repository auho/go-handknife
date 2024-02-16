package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

func LoadMysql(name string, dbname string, mc MysqlConf) (*gorm.DB, error) {
	if mc.Host == "" {
		return nil, errors.New(fmt.Sprintf("msyql[%s] host is empty", name))
	}

	if mc.Port == "" {
		mc.Port = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", mc.Username, mc.Password, mc.Host, mc.Port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold: time.Second,  // 慢 SQL 阈值
				LogLevel:      logger.Error, // 日志级别
			},
		)})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("msyql[%s] %v", name, err))
	}

	err = db.Select("select version()").Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("msyql[%s] connection %v", name, err))
	}

	return db, nil
}
