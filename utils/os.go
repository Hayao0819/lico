package utils

import (
	"errors"
	"os"
	"runtime"
)

// OS依存の情報を保持します
type osEnv struct{
	Home string
}

// 新しいosEnvを生成します
func newOSEnv ()(osEnv){
	env:= new(osEnv)
	homedir, _ := os.UserHomeDir()
	env.Home = homedir
	return *env
}

// Linux特有のディレクトリ情報
var LinuxDirs osEnv = func ()(osEnv)  {
	return newOSEnv()
}()


// Darwin特有のディレクトリ情報
var MacDirs osEnv = func ()(osEnv)  {
	return newOSEnv()
}()

// Windows特有のディレクトリ情報
var WindowsDirs osEnv = func()(osEnv){
	return newOSEnv()
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


