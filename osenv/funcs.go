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
