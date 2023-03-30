package osenv

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
