package osenv

// LinuxのOS情報
var LinuxInfo E = func() E {
	env := newOSEnv()
	env["OS"] = "linux"
	return env
}()

// DarwinのOS情報
var MacInfo E = func() E {
	env := newOSEnv()
	env["OS"] = "darwin"
	return env
}()

// WindowsのOS情報
var WindowsInfo E = func() E {
	env := newOSEnv()
	env["OS"] = "windows"
	return env
}()
