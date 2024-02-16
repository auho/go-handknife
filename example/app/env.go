package app

import "github.com/auho/go-handknife/emergencybox/app"

const dev = "dev"
const prod = "prod"

type Environment struct {
	Env string `yaml:"env"`

	// custom
	BaseRedis app.RedisConf `yaml:"base-redis"`
	BaseDB    app.MysqlConf `yaml:"base-db"`
	BaseEs    app.EsConf    `yaml:"base-es"`
	BaseKafka app.KafkaConf `yaml:"base-kafka"`
}

func (e Environment) IsDev() bool {
	return e.Env == dev
}

func (e Environment) IsProd() bool {
	return e.Env == prod
}
