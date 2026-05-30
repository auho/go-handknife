package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/auho/go-handknife/blade/app"
	"github.com/auho/go-handknife/blade/suites"
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
	client, err := App.GetBaseRedis(context.Background())
	if err != nil {
		panic(fmt.Errorf("UseCase.BaseRedis: %w", err))
	}

	return client
}

func (uc *UseCase) BaseMysql() *gorm.DB {
	db, err := App.GetBaseDB(context.Background())
	if err != nil {
		panic(fmt.Errorf("UseCase.BaseDB: %w", err))
	}
	return db
}

func (uc *UseCase) BaseEs() *elasticsearch.Client {
	es, err := App.GetBaseEs(context.Background())
	if err != nil {
		panic(fmt.Errorf("UseCase.BaseEs: %w", err))
	}

	return es
}

func (uc *UseCase) App() *Application {
	return App
}
