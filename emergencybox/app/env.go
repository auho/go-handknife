package app

type MysqlConf struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisConf struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Auth string `yaml:"auth"`
	DB   int    `yaml:"db"`
}

type EsConf struct {
	Name     string   `yaml:"name"`
	Address  []string `yaml:"address"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type KafkaConf struct {
	Name    string `yaml:"name"`
	NetWork string `yaml:"network"`
	Address string `yaml:"address"`
}
