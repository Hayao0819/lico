package osenv_test

import (
	"testing"

	"github.com/Hayao0819/lico/osenv"
)

func env_match_string(t *testing.T , e osenv.E ,key string, correct string){
	current := e.Get(key)
	if current != correct{
		t.Errorf("the value of %s is not %s. corrent value is %s", key, correct, current)
	}
}

func TestOS(t *testing.T){
	type test struct{
		name string
		env osenv.E
	}

	tests := []test{
		{
			name: "linux",
			env:  osenv.LinuxInfo,
		},
		{
			name: "windows",
			env: osenv.WindowsInfo,
		},
		{
			name: "darwin",
			env: osenv.MacInfo,
		},
	}

	for _ ,tt := range tests{

		switch tt.name {
			case "linux":
				t.Run(tt.name, func(t *testing.T) {
					env_match_string(t, tt.env, "OS", "linux")
				})
			case "windows":
				t.Run(tt.name, func(t *testing.T) {
					env_match_string(t, tt.env, "OS", "windows")
				})
			case "darwin":
				t.Run(tt.name, func(t *testing.T) {
					env_match_string(t, tt.env, "OS", "darwin")
				})
		}
	}
}

//func Test
