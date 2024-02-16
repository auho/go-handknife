package toolkit

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/auho/go-handknife/emergencybox/toolkit/convert"

	"github.com/go-redis/redis/v8"
)

// OutputArrayMap
// format []string{"2%-30s", "4%s"}
// "2%-30s" map 的第 2 列，左对齐 30 宽度
// "4%s" 	map 的第 4 列
func OutputArrayMap(ctx context.Context, c *redis.Client, key string, format []string) (string, error) {
	s, err := c.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	_cmd := `cut -b 4- | tr '[]{}' ' ' | sed 's/ , /\n/g' | tr '"' ' ' | awk -F ' : | , | :|, ' '%s'`
	_cmd = fmt.Sprintf(_cmd, printfFormant(format))

	_c := exec.Command("bash", "-c", _cmd)
	var _out strings.Builder
	_c.Stdin = strings.NewReader(s)
	_c.Stdout = &_out
	err = _c.Run()
	if err != nil {
		return "", err
	}

	return convert.ContainsUnicodeToZh(_out.String()), nil
}

// OutputArrayArray
// format []string{"2%-30s", "4%s"}
// "2%-30s" map 的第 2 列，左对齐 30 宽度
// "4%s" 	map 的第 4 列
func OutputArrayArray(ctx context.Context, c *redis.Client, key string, format []string) (string, error) {
	s, err := c.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	_cmd := `cut -b 2- | tr '}' '\n' | tr '[{,]' '    ' | tr  -d '"' | tr -s ' ' | awk '%s'`
	_cmd = fmt.Sprintf(_cmd, printfFormant(format))

	_c := exec.Command("bash", "-c", _cmd)
	var _out strings.Builder
	_c.Stdin = strings.NewReader(s)
	_c.Stdout = &_out
	err = _c.Run()
	if err != nil {
		log.Fatal(err)
	}

	return convert.ContainsUnicodeToZh(_out.String()), nil
}

func printfFormant(ss []string) string {
	s := ""
	for _, v := range ss {
		_v := strings.Split(v, "%")
		s = s + `printf "%` + _v[1] + `", $` + _v[0] + ";"
	}

	return "{" + s + `print ""}`
}
