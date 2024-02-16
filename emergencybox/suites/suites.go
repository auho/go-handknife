package suites

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/auho/go-handknife/emergencybox/suites/output"
	"github.com/auho/go-handknife/emergencybox/suites/template"
	output2 "github.com/auho/go-toolkit/console/output"
	"github.com/auho/go-toolkit/farmtools/convert/types/structs/maps"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const nameKey = "key"
const nameValue = "value"
const nameMember = "member"
const nameScore = "score"

type Suite struct {
	Cmd       *cobra.Command
	Validator *validator.Validate

	pfFunc  map[int][]func() string // map[pf func index][]func
	pfIndex int

	mutex sync.Mutex
}

func (se *Suite) Init(cmd *cobra.Command) {
	se.Cmd = cmd
	se.Validator = validator.New()
	se.pfFunc = make(map[int][]func() string)
}

func (se *Suite) CmdVisit(cmd *cobra.Command) string {
	var args []string
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		value := flag.Value.String()
		args = append(args, fmt.Sprintf("--%s='%s'", flag.Name, value))
	})

	return strings.Join(args, " ")
}

func (se *Suite) PfSliceMapStringAny(title string, fn func() ([]map[string]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		sm, p, err := fn()
		if err == nil {
			sm = se.renderSliceMap(sm, p)
		}

		return sm, p, err
	})
}

func (se *Suite) PfMapStringString(title string, fn func() (map[string]string, *Post, error)) {
	se.pfIndexIn()

	se.PfMapStringAny(title, func() (map[string]any, *Post, error) {
		newM := make(map[string]any)
		m, p, err := fn()
		if err == nil {
			for _k, _v := range m {
				newM[_k] = _v
			}
		}

		return newM, p, err
	})
}

func (se *Suite) PfMapStringAny(title string, fn func() (map[string]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		m, p, err := fn()
		if err == nil {
			var keys []string
			for _k := range m {
				keys = append(keys, _k)
			}

			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})

			m = se.renderMap(m, p)

			for _, k := range keys {
				sm = append(sm, map[string]any{
					"key":   k,
					"value": m[k],
				})
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfSliceStruct(title string, fn func() ([]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		sa, p, err := fn()
		if err == nil {
			headHasNotNil := true
			for _, _s := range sa {
				if headHasNotNil && _s == nil {
					continue
				}

				headHasNotNil = false

				_m, err1 := maps.MapStringAnyFromStruct(_s)
				if err1 != nil {
					return nil, p, err1
				}

				_m = se.renderMap(_m, p)

				sm = append(sm, _m)
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfSliceStructToKv(title string, fn func() ([]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		_sa, p, err := fn()
		if err == nil && len(_sa) > 0 {
			var sa []any
			headHasNotNil := true
			for _, _a := range _sa {
				if headHasNotNil && _a == nil {
					continue
				}

				headHasNotNil = false

				sa = append(sa, _a)
			}

			_m, err1 := maps.MapStringAnyFromStruct(sa[0])
			if err1 != nil {
				return nil, p, err1
			}

			var keys []string
			for _k := range _m {
				keys = append(keys, _k)
			}

			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})

			for _, _s := range sa {
				_tm, err2 := maps.MapStringAnyFromStruct(_s)
				if err2 != nil {
					return nil, p, err1
				}

				_tm = se.renderMap(_tm, p)

				for _, key := range keys {
					sm = append(sm, map[string]any{
						nameKey:   key,
						nameValue: _tm[key],
					})
				}

				sm = append(sm, nil)
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfStruct(title string, fn func() (any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		a, p, err := fn()
		if err == nil {
			_m, err1 := maps.MapStringAnyFromStruct(a)
			if err1 != nil {
				return nil, p, err1
			}

			var keys []string
			for _k := range _m {
				keys = append(keys, _k)
			}

			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})

			_m = se.renderMap(_m, p)

			for _, _k := range keys {
				sm = append(sm, map[string]any{
					nameKey:   _k,
					nameValue: _m[_k],
				})
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfSlice(title string, fn func() ([]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		m, p, err := fn()
		if err == nil {
			for _, _s := range m {
				sm = append(sm, map[string]any{
					nameValue + ":": _s,
				})
			}

			sm = se.renderSliceMap(sm, p)
		}

		if p == nil {
			p = NewPost()
		}

		p.AddTemplate(func(t *template.Template) {
			t.AddOutputSetting(func(writer table.Writer) {
				writer.SetStyle(table.Style{
					Format: table.FormatOptions{
						Header: text.FormatUpper,
					},
					Options: table.Options{
						DrawBorder:     false,
						SeparateHeader: true,
					},
				})

				writer.SetAutoIndex(false)
			})
		})

		return sm, p, err
	})
}

// PfSliceToKV
// odd as key
// even as value
//
// key
// value
func (se *Suite) PfSliceToKV(title string, fn func() ([]any, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		m, p, err := fn()
		if err == nil {
			_len := len(m)
			for i := 0; i < _len; i++ {
				sm = append(sm, map[string]any{
					nameKey:   m[i],
					nameValue: m[i+1],
				})
			}

			sm = se.renderSliceMap(sm, p)
		}

		return sm, p, err
	})
}

// PfRedisSliceZ
// No
// member
// score
func (se *Suite) PfRedisSliceZ(title string, fn func() ([]redis.Z, *Post, error)) {
	se.pfIndexIn()

	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		rz, p, err := fn()
		if err == nil {
			for _, z := range rz {
				sm = append(sm, map[string]any{
					nameMember: z.Member,
					nameScore:  strconv.FormatFloat(z.Score, 'f', -1, 64),
				})
			}

			sm = se.renderSliceMap(sm, p)
		}

		return sm, p, err
	})
}

func (se *Suite) printSliceMapStringAny(title string, fn func() ([]map[string]any, *Post, error)) {
	se.printFunc(title, func() (*Post, string, error) {
		var body string
		sm, p, err := fn()
		if err == nil {
			body, err = se.renderSliceMapStringAny(sm, p)
		}

		return p, body, err
	})
}

func (se *Suite) printFunc(title string, fn func() (*Post, string, error)) {
	s := &strings.Builder{}
	s.WriteString(se.stringTitle(title) + "\n")

	p, body, err := fn()

	se.pfIndexExec(func(_f func() string) {
		s.WriteString(_f())
		s.WriteString("\n")
	})

	if err != nil {
		se.PrintlnErr("[Error] ", err)
	} else {
		s.WriteString(body)
	}

	se.Cmd.Print(s.String())
	se.Cmd.Println("")

	if p == nil {
		p = NewPost()
	}

	extra := p.Extra
	if extra != nil {
		err = p.Extra()
		if err != nil {
			var nErr validator.ValidationErrors
			switch {
			case errors.As(err, &nErr):
				for k, fieldError := range nErr {
					se.PrintlnErr("extra:", k, fieldError)
				}
			default:
				se.PrintlnErr("extra:", err)
			}
		}

		se.Cmd.Println("")
	}
}

func (se *Suite) PfFunc(title string, fn func() (any, error)) {
	se.pfIndexIn()

	se.PrintlnTitle(title)
	v, err := fn()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			v = "redis result is nil"
		}
	}

	se.Cmd.Printf(" -> %v\n\n", v)
}

func (se *Suite) PfVoid(title string, fn func() error) {
	se.pfIndexIn()

	se.PrintlnTitle(title)

	err := fn()

	se.pfIndexExec(func(_f func() string) {
		se.Cmd.Println(_f())
	})

	if err != nil {
		se.PrintlnErr("[Error] ", err)
	}
}

func (se *Suite) pfIndexIn() {
	se.mutex.Lock()
	se.pfIndex++
	se.mutex.Unlock()
}

func (se *Suite) pfIndexOut() {
	se.mutex.Lock()
	se.pfIndex--
	se.mutex.Unlock()
}

func (se *Suite) pfIndexExec(fn func(func() string)) {
	for _, _fn := range se.pfFunc[se.pfIndex] {
		fn(_fn)
	}

	delete(se.pfFunc, se.pfIndex)
	se.pfIndexOut()
}

func (se *Suite) PfBody(body string) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return se.stringBody(body)
	})
}

func (se *Suite) PfLn(a ...any) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return fmt.Sprint(a...)
	})
}

func (se *Suite) PfErr(a ...any) {
	se.pfFunc[se.pfIndex] = append(se.pfFunc[se.pfIndex], func() string {
		return se.stringErr(a...)
	})
}

func (se *Suite) Println(s string) {
	se.Cmd.Println(s)
}

func (se *Suite) PrintlnAny(i ...any) {
	se.Cmd.Println(i...)
}

func (se *Suite) Watch(title string, totalDuration, interval time.Duration, fn func() ([]string, error)) {
	se.PrintlnTitle(title)

	_r, cancel := output2.NewRefreshWithCancel(
		output2.WithContent(fn),
		output2.WithInterval(interval),
	)

	_r.Start()

	<-time.After(totalDuration)
	cancel()
}

func (se *Suite) Func(title string, fn func() error) {
	se.PrintlnTitle(title)
	err := fn()
	if err != nil {
		se.PrintlnErr(err)
	}

	se.Cmd.Println()
}

func (se *Suite) PrintlnTitle(title string) {
	se.Cmd.Println(se.stringTitle(title))
}

func (se *Suite) PrintlnBody(body string) {
	se.Cmd.Println(se.stringBody(body))
}

func (se *Suite) PrintlnErr(a ...any) {
	se.Cmd.PrintErrln(se.stringErr(a...))
}

func (se *Suite) renderSliceMap(sm []map[string]any, p *Post) []map[string]any {
	if p == nil || p.Template == nil {
		return sm
	}

	var nsm []map[string]any
	for _, m := range sm {
		nsm = append(nsm, se.renderMap(m, p))
	}

	return nsm
}

func (se *Suite) renderMap(m map[string]any, p *Post) map[string]any {
	if p == nil || p.Template == nil {
		return m
	}

	nm := make(map[string]any, len(m))
	for k, v := range m {
		nm[k] = p.Template.FieldExec(k, v)
	}

	return nm
}

func (se *Suite) renderSliceMapStringAny(sm []map[string]any, p *Post) (string, error) {
	if p == nil {
		p = NewPost()
	}

	_output := output.NewOutput()
	body, err := _output.Output(sm, p.Template, func(writer table.Writer) {
		writer.SetAutoIndex(true)
		for _, fn := range p.Template.OutputSettings() {
			fn(writer)
		}
	})
	if err != nil {
		return "", err
	}

	return body, nil
}

func (se *Suite) stringTitle(title string) string {
	return text.FgBlue.Sprint(title)
}

func (se *Suite) stringBody(body string) string {
	return text.FgYellow.Sprint(body)
}

func (se *Suite) stringErr(a ...any) string {
	return text.FgHiRed.Sprint(a...)
}
