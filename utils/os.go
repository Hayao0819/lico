package utils

import (
	"errors"
	"os"
	"runtime"
)

type osEnv struct{
	Home string
}

func newOSEnv ()(osEnv){
	env:= new(osEnv)
	homedir, _ := os.UserHomeDir()
	env.Home = homedir
	return *env
}

var LinuxDirs osEnv = func ()(osEnv)  {
	return newOSEnv()
}()

var MacDirs osEnv = func ()(osEnv)  {
	return newOSEnv()
}()

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


