package utils

import (
	"errors"
	//"fmt"
	"os"
	"os/user"
	"runtime"
	"sort"
	"strings"

	"github.com/Hayao0819/lico/vars"
)

// OS依存の情報を保持します
type osEnv map[string]string

// 新しいosEnvを生成します
func newOSEnv() osEnv {
	user, _ := user.Current()
	env := map[string]string{
		"Home":     GetHomeDir(),
		"OS":       "",
		"UserName": user.Username,
	}

	for index, value := range getEnvVars() {
		env[index] = value
	}

	return osEnv(env)
}

func getEnvVars() map[string]string {
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

// Linux特有のディレクトリ情報
var LinuxDirs osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "linux"
	return env
}()

// Darwin特有のディレクトリ情報
var MacDirs osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "darwin"
	return env
}()

// Windows特有のディレクトリ情報
var WindowsDirs osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "windows"
	return env
}()

func GetOSEnv() (osEnv, error) {
	var env osEnv
	switch runtime.GOOS {
	case "windows":
		env = WindowsDirs
	case "linux":
		env = LinuxDirs
	case "darwin":
		env = MacDirs
	default:
		return env, errors.New("unsupported os")
	}

	env.Add("Repo", vars.RepoDir)
	env.Add("List", vars.BaseListFile)
	return env, nil
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
