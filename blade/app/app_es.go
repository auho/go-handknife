package app

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
)

type EsConf struct {
	Name       string   `yaml:"name"`
	Address    []string `yaml:"address"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	MaxRetries int      `yaml:"max_retries"`
}

func (a *Application) GetEs(ctx context.Context, conf EsConf) (*elasticsearch.Client, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if e, ok := a.esHub[conf.Name]; ok {
		return e, nil
	}

	e, err := a.openEs(ctx, conf)
	if err != nil {
		return nil, err
	}

	a.esHub[conf.Name] = e

	return e, nil
}

func (a *Application) openEs(_ context.Context, conf EsConf) (*elasticsearch.Client, error) {
	if len(conf.Address) <= 0 {
		return nil, fmt.Errorf("elasticSearch[%s] address is empty", conf.Name)
	}

	maxRetries := conf.MaxRetries
	if maxRetries <= 0 {
		maxRetries = 3
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:  conf.Address,
		Username:   conf.Username,
		Password:   conf.Password,
		MaxRetries: maxRetries,
	})
	if err != nil {
		return nil, fmt.Errorf("elasticSearch[%s] client %v", conf.Name, err)
	}

	resp, err := es.Ping()
	if err != nil {
		return nil, fmt.Errorf("elasticSearch[%s] ping error %v", conf.Name, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("elasticSearch[%s] ping error %s", conf.Name, resp.Status())
	}

	return es, nil
}
