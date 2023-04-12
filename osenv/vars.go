package osenv

// LinuxのOS情報
var LinuxInfo osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "linux"
	return env
}()

// DarwinのOS情報
var MacInfo osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "darwin"
	return env
}()

// WindowsのOS情報
var WindowsInfo osEnv = func() osEnv {
	env := newOSEnv()
	env["OS"] = "windows"
	return env
}()
