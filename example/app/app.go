package app

import (
	"context"

	"github.com/auho/go-handknife/blade/app"
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

func (a *Application) GetBaseDB(ctx context.Context) (*gorm.DB, error) {
	return a.app.GetDB(ctx, a.env.BaseDB)
}

func (a *Application) GetBaseRedis(ctx context.Context) (*redis.Client, error) {
	return a.app.GetRedis(ctx, a.env.BaseRedis)
}

func (a *Application) GetBaseEs(ctx context.Context) (*elasticsearch.Client, error) {
	return a.app.GetEs(ctx, a.env.BaseEs)
}

func (a *Application) GetEnv() Environment {
	return a.env
}
