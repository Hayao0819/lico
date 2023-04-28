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

func TestIsFile(t *testing.T){
	args := []struct{
		path string
		expect bool
	}{
		{
			path: path.Join(vars.RepoDir , "/config"),
			expect: false,
		},
		{
			path: path.Join(vars.RepoDir , "README.md"),
			expect: true,
		},
	}
	for _, arg := range args {
		if utils.IsFile(arg.path) != arg.expect {
			t.Errorf("utils.IsFile(%s) = %t", arg.path, !arg.expect)
		}
	}
}

func TestIsEmpty(t *testing.T){
	args := []struct{
		str string
		expect bool
	}{
		{
			str: "",
			expect: true,
		},
		{
			str: " ",
			expect: true,
		},
		{
			str: "a",
			expect: false,
		},
	}
	for _, arg := range args {
		if utils.IsEmpty(arg.str) != arg.expect {
			t.Errorf("utils.IsEmpty(%s) = %t", arg.str, !arg.expect)
		}
	}
}

// あとで実装する
/*
func TestIsSymlink(t *testing.T){

}
*/


func TestExists(t *testing.T){
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
			expect: true,
		},
		{
			path: path.Join(vars.RepoDir , "not_exist"),
			expect: false,
		},
	}
	for _, arg := range args {
		if utils.Exists(arg.path) != arg.expect {
			t.Errorf("utils.Exists(%s) = %t", arg.path, !arg.expect)
		}
	}
}


func TestCommandExists(t *testing.T){
	args := []struct{
		command string
		expect bool
	}{
		{
			command: "ls",
			expect: true,
		},
		{
			command: "go",
			expect: true,
		},
		{
			command: "not_exist",
			expect: false,
		},
	}
	for _, arg := range args {
		if utils.CommandExists(arg.command) != arg.expect {
			t.Errorf("utils.CommandExists(%s) = %t", arg.command, !arg.expect)
		}
	}
}
