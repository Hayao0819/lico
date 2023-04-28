package utils_test

import (
	"path"
	"testing"

	"github.com/Hayao0819/lico/tester"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)


func TestMain(m *testing.M){
	tester.CommonTestMain("../example")(m)
}

func TestIsDir(t *testing.T) {
	args := []struct{
		path string
		expect bool
	}{
		{
			path: path.Join(vars.RepoDir , "/config"),
			expect: true,
		},
		{
			path: path.Join(vars.RepoDir , "README.md"),
			expect: false,
		},
	}
	for _, arg := range args {
		if utils.IsDir(arg.path) != arg.expect {
			t.Errorf("utils.IsDir(%s) = %t", arg.path, !arg.expect)
		}
	}
}
