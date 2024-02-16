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

// KeysVerboseExecute
// keys verbose execute
// 多个关键字 Verbose 的执行
type KeysVerboseExecute struct {
	ksv map[string]Verbose

	// 匹配 keys 的 regex
	regexps map[string]*regexp.Regexp
}

// NewKeysVerboseExecute
// 多个关键字的 Verbose
// keys => Verbose
// /regex/ => Verbose
func NewKeysVerboseExecute(kv map[string]Verbose) *KeysVerboseExecute {
	regexps := make(map[string]*regexp.Regexp)
	for k := range kv {
		if k[0:1] == "/" && k[len(k)-1:] == "/" {
			regexps[k] = regexp.MustCompile(k[1 : len(k)-1])
		}
	}

	return &KeysVerboseExecute{
		ksv:     kv,
		regexps: regexps,
	}
}

// ExecuteKv
// 查找 key 的所有 verbose，传入 value 执行
func (k *KeysVerboseExecute) ExecuteKv(key string, value any) any {
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

	return k.Execute(value, kv)
}

// Execute
// 执行指定 value 和 verbose
func (k *KeysVerboseExecute) Execute(v any, kv Verbose) any {
	for _, kfn := range kv {
		v = kfn(v)
	}

	return v
}
