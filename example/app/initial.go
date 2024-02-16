package app

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Initialization() {
	_f := "conf/config.yaml"
	b, err := os.ReadFile(_f)
	if err != nil {
		log.Fatal("read config file error:", err)
	}

	var env Environment
	err = yaml.Unmarshal(b, &env)
	if err != nil {
		log.Fatal("unmarshal config file error:", err)
	}

	App = &Application{}
	err = App.init(env)
	if err != nil {
		log.Fatal("init application error:", err)
	}
}
