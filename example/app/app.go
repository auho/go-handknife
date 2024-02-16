package app

import (
	"github.com/auho/go-handknife/emergencybox/app"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var App *Application

type Application struct {
	app *app.Application

	env Environment
}

func (a *Application) init(env Environment) error {
	a.app = app.NewApplication()
	a.env = env

	return nil
}

func (a *Application) GetBaseRedis() *redis.Client {
	return a.app.GetRedis(a.env.BaseRedis)
}

func (a *Application) GetBaseDB() *gorm.DB {
	return a.app.GetDB(a.env.BaseDB)
}

func (a *Application) GetBaseEs() *elasticsearch.Client {
	return a.app.GetEs(a.env.BaseEs)
}

func (a *Application) GetEnv() Environment {
	return a.env
}
