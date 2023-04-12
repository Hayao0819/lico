package osenv

import (
	"errors"
	"runtime"

	"github.com/Hayao0819/lico/vars"
)

func Get() (osEnv, error) {
	var env osEnv
	switch runtime.GOOS {
	case "windows":
		env = WindowsInfo
	case "linux":
		env = LinuxInfo
	case "darwin":
		env = MacInfo
	default:
		return env, errors.New("unsupported os")
	}

	env.Add("Repo", vars.RepoDir)
	env.Add("List", vars.BaseListFile)
	return env, nil
}
