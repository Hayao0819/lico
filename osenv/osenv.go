package osenv

import (
	"os"
	"os/user"
	"sort"
	"strings"

	"github.com/Hayao0819/lico/utils"
)

// OS依存の情報を保持します
type osEnv map[string]string

// 新しいosEnvを生成します
func newOSEnv() osEnv {
	user, _ := user.Current()
	env := map[string]string{
		"Home":     utils.GetHomeDir(),
		"OS":       "",
		"UserName": user.Username,
	}

	for index, value := range getVars() {
		env[index] = value
	}

	return osEnv(env)
}

// 環境変数を取得
func getVars() map[string]string {
	rtn := map[string]string{}
	for _, envS := range os.Environ() {
		env := strings.Split(envS, "=")
		if strings.HasPrefix(env[0], "LICO_") {
			index := strings.TrimPrefix(env[0], "LICO_")
			rtn[index] = env[1]
		}
	}
	return rtn
}

func (o *osEnv) Add(key, value string) *osEnv {
	m := *o
	m[key] = value
	return (*osEnv)(&m)
}

func (o *osEnv) Get(key string) string {
	m := *o
	return m[key]
}

func (env *osEnv) GetKeys() []string {
	var arr []string
	for index := range *env {
		arr = append(arr, index)
	}

	return arr
}

func (env osEnv) GetSortedKeys() []string {
	type osEvnStruct struct {
		name  string
		value string
	}

	var envS []osEvnStruct
	for _, key := range env.GetKeys() {
		envS = append(envS, osEvnStruct{name: key, value: env[key]})
	}

	sort.Slice(envS, func(i, j int) bool {
		return len(envS[i].value) > len(envS[j].value)
	})

	var rtn []string
	for _, key := range envS {
		rtn = append(rtn, key.name)
	}
	return rtn
}
