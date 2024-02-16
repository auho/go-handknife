package sh

import (
	"strings"

	"github.com/auho/go-handknife/env/form/config"
)

func parseRedisConf(c string) config.RedisConf {
	rc := config.RedisConf{}

	argsString := strings.Split(c, " ")
	for i := 0; i < len(argsString); i++ {
		_flag := strings.TrimLeft(argsString[i], "-")
		_arg := argsString[i+1]
		switch _flag {
		case "h":
			rc.Addr = _arg
		case "p":
			rc.Port = _arg
		case "a":
			rc.Password = _arg
		default:

		}

		i++
	}

	return rc
}
