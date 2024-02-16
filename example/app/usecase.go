package app

import (
	"sync"

	"github.com/auho/go-handknife/emergencybox/app"
	"github.com/auho/go-handknife/emergencybox/suites"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type UseCase struct {
	suites.Suite
	app.UseCase

	mutex sync.Mutex
}

func (uc *UseCase) InitUseCase(c *cobra.Command) {
	uc.Init(c)
}

func (uc *UseCase) BaseRedis() *redis.Client {
	return App.GetBaseRedis()
}

func (uc *UseCase) BaseMysql() *gorm.DB {
	return App.GetBaseDB()
}

func (uc *UseCase) BaseEs() *elasticsearch.Client {
	return App.GetBaseEs()
}

func (uc *UseCase) App() *Application {
	return App
}
