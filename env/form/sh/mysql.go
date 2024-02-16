package sh

import (
	"strings"

	"github.com/auho/go-handknife/env/form/config"
)

func parseMysqlConf(c string) config.MysqlConf {
	mc := config.MysqlConf{}

	argsString := strings.Split(c, " ")
	for i := 0; i < len(argsString); i++ {
		_flag := strings.TrimLeft(argsString[i], "-")

		switch _flag {
		case "h", "P", "u":
			_arg := argsString[i+1]
			switch _flag {
			case "h":
				mc.Host = _arg
			case "P":
				mc.Port = _arg
			case "u":
				mc.Username = _arg
			}
		default:
			_first := _flag[0:1]
			switch _first {
			case "p":
				mc.Password = _flag[1:]
			}
		}

		i++
	}

	return mc
}
