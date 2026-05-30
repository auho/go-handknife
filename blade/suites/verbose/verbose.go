package verbose

import (
	"regexp"
)

// Func
// function for verbose value
type Func func(any) any

// Verbose
// functions for verbose value
// 详细（格式化、转化等）字段内容
type Verbose []Func

func NewVerbose(fs ...Func) Verbose {
	return fs
}

// Keys
// keys verbose execute
// 多个关键字 Verbose 的执行
type Keys struct {
	ksv map[string]Verbose // map[key]verbose

	// 匹配 keys 的 regex
	regexps map[string]*regexp.Regexp // map[key]*regexp.Regexp
}

// NewKeys
// 多个关键字的 Verbose
// keys => Verbose
// /regex/ => Verbose
func NewKeys() *Keys {
	return &Keys{
		ksv:     make(map[string]Verbose),
		regexps: make(map[string]*regexp.Regexp),
	}
}

func NewKeysWithVerbose(mv map[string]Verbose) *Keys {
	k := NewKeys()
	k.AddWithMap(mv)

	return k
}

func (k *Keys) Add(key string, vb Verbose) *Keys {
	k.ksv[key] = vb

	if key[0:1] == "/" && key[len(key)-1:] == "/" {
		k.regexps[key] = regexp.MustCompile(key[1 : len(key)-1])
	}

	return k
}

func (k *Keys) AddWithMap(mv map[string]Verbose) *Keys {
	for key, v := range mv {
		k.Add(key, v)
	}

	return k
}

// ExecuteKey
// 查找 key 的所有 verbose，传入 value 执行
func (k *Keys) ExecuteKey(key string, value any) any {
	// 字符串匹配 keys
	kv, ok := k.ksv[key]
	if !ok {
		// regex 匹配 keys
		for _key, r := range k.regexps {
			if r.MatchString(key) {
				kv = k.ksv[_key]
			}
		}

		// 使用 default
		if len(kv) <= 0 {
			if kv, ok = k.ksv[""]; !ok {
				return value
			}
		}
	}

	return k.execute(value, kv)
}

// execute
// 执行指定 value 和 verbose
func (k *Keys) execute(v any, kv Verbose) any {
	for _, kfn := range kv {
		v = kfn(v)
	}

	return v
}
