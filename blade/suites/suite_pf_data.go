package suites

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/auho/go-toolkit/farmtools/convert/types/structs/maps"
	"github.com/go-redis/redis/v8"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func (se *Suite) PfSliceMapStringAny(title string, fn func() ([]map[string]any, *Post, error)) {
	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		return fn()
	})
}

func (se *Suite) PfMapStringString(title string, fn func() (map[string]string, *Post, error)) {
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

				sm = append(sm, _m)
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfSliceStructToKv(title string, fn func() ([]any, *Post, error)) {
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
	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		m, p, err := fn()
		if err == nil {
			for _, _s := range m {
				sm = append(sm, map[string]any{
					nameValue + ":": _s,
				})
			}
		}

		if p == nil {
			p = NewPost()
		}

		p.SettingOutput(func(writer table.Writer) {
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

		return sm, p, err
	})
}

func (se *Suite) PfSliceToKV(title string, fn func() ([]any, *Post, error)) {
	se.printSliceMapStringAny(title, func() ([]map[string]any, *Post, error) {
		var sm []map[string]any
		m, p, err := fn()
		if err == nil {
			_len := len(m)
			if _len%2 == 1 {
				_len += 1
				m = append(m, "")
			}

			for i := 0; i+1 < _len; i += 2 {
				sm = append(sm, map[string]any{
					nameKey:   m[i],
					nameValue: m[i+1],
				})
			}
		}

		return sm, p, err
	})
}

func (se *Suite) PfRedisSliceZ(title string, fn func() ([]redis.Z, *Post, error)) {
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
		}

		return sm, p, err
	})
}

func (se *Suite) PfFunc(title string, fn func() (any, error)) {
	se.pfIndexIn()

	se.PrintlnTitle(title)
	v, err := fn()

	se.pfIndexExec(func(_f func() string) {
		se.Cmd.Println(_f())
	})

	if err != nil {
		if errors.Is(err, redis.Nil) {
			v = "redis result is nil"
		}
	}

	se.Println(fmt.Sprintf(" -> %v\n\n", v))
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
