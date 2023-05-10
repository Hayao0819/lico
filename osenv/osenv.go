package osenv

import (
	"os"
	"os/user"
	"runtime"
	"sort"
	"strings"

	"github.com/Hayao0819/go-distro"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)

// OS依存の情報を保持します
type E map[string]string

// 新しいosEnvを生成します
func Get() (E, error) {
	user, _ := user.Current()
	d := distro.Get()
	env := map[string]string{
		"Home":     utils.GetHomeDir(),
		"GOOS":     runtime.GOOS,
		"OS":       d.Name(),
		"OSVer":    d.Version().ID(),
		"UserName": user.Username,
		"Repo":     vars.GetRepoDir(),
		"List":     vars.GetList(),
	}

	for index, value := range getVars() {
		env[index] = value
	}

	return E(env), nil
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

func (o *E) Add(key, value string) *E {
	m := *o
	m[key] = value
	return (*E)(&m)
}

func (o *E) Get(key string) string {
	m := *o
	return m[key]
}

func (env *E) GetKeys() []string {
	var arr []string
	for index := range *env {
		arr = append(arr, index)
	}

	return arr
}

func (env E) GetSortedKeys() []string {
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
