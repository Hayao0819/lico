package utils

import (
	"errors"
	"os"
	"runtime"
	"os/user"
)

// OS依存の情報を保持します
type osEnv map[string]string

// 新しいosEnvを生成します
func newOSEnv ()(osEnv){
	homedir, _ := os.UserHomeDir()
	user, _ := user.Current()
	env:= osEnv{"Home": homedir, "OS": "", "UserName": user.Username}
	return env
}

// Linux特有のディレクトリ情報
var LinuxDirs osEnv = func ()(osEnv)  {
	env := newOSEnv()
	env["OS"]="linux"
	return env
}()


// Darwin特有のディレクトリ情報
var MacDirs osEnv = func ()(osEnv)  {
	env := newOSEnv()
	env["OS"]="darwin"
	return env
}()

// Windows特有のディレクトリ情報
var WindowsDirs osEnv = func()(osEnv){
	env := newOSEnv()
	env["OS"]="windows"
	return env
}()

func GetOSEnv()(osEnv, error){
	switch runtime.GOOS{
		case "windows":
			return WindowsDirs, nil
		case "linux":
			return LinuxDirs,nil
		case "darwin":
			return MacDirs, nil
		default:
			return newOSEnv(), errors.New("unsupported os")
	}
}


