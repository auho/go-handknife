package sh

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/auho/go-handknife/env/form"
	config2 "github.com/auho/go-handknife/env/form/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var _ form.Former = (*Sh)(nil)

type Sh struct {
	content      []byte
	configString string
}

func NewSh(filePath string) (*Sh, error) {
	cb, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &Sh{
			content:      cb,
			configString: string(cb),
		},
		nil
}

func (s *Sh) LoadEnv() (string, error) {
	_env, err := s.extractMatchContent(`(?m)env\=(.+)`, s.configString)
	if err != nil {
		return "", fmt.Errorf("load env;%w", err)
	}

	return _env, nil
}

func (s *Sh) LoadMysql(name string, dbname string) (*gorm.DB, error) {
	_name := strings.ToUpper(name[0:1]) + name[1:]
	c, err := s.extractMatchContent(fmt.Sprintf(`(?m)mysql%s="([^"]+)"`, _name), s.configString)
	if err != nil {
		panic(fmt.Errorf("loadMysql %w", err))
	}

	mc := parseMysqlConf(c)

	return config2.LoadMysql(name, dbname, mc)
}

func (s *Sh) LoadRedis(name string) (*redis.Client, error) {
	_name := strings.ToUpper(name[0:1]) + name[1:]
	c, err := s.extractMatchContent(fmt.Sprintf(`(?m)redis%s="([^"]+)"`, _name), s.configString)
	if err != nil {
		panic(fmt.Errorf("loadRedis %w", err))
	}

	rc := parseRedisConf(c)

	return config2.LoadRedis(name, rc)
}

func (s *Sh) LoadServer(name string) (config2.Server, error) {
	ip, err := s.extractMatchContent(`(?m)serverIp="([^"]+)"`, s.configString)
	if err != nil {

	}

	port, err := s.extractMatchContent(`(?m)serverPort="([^"]+)"`, s.configString)
	if err != nil {

	}

	token, err := s.extractMatchContent(`(?m)serverToken="([^"]+)"`, s.configString)
	if err != nil {

	}

	return config2.Server{
		Ip:    ip,
		Port:  port,
		Token: token,
	}, nil
}

func (s *Sh) extractMatchContent(re string, ss string) (string, error) {
	res := regexp.MustCompile(re).FindStringSubmatch(ss)
	if len(res) < 2 {
		return "", fmt.Errorf("%s not found", re)
	}

	return res[1], nil
}
