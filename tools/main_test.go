package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Hayao0819/lico/tester"
)

func TestArtifactCmd(t *testing.T) {

	testcase := []struct{
		json string
		path string
	}{
		{
			json: `[{"name":"lico","path":"/Users/user/Git/lico/dist/lico-build_darwin_arm64/lico","goos":"darwin","goarch":"arm64","internal_type":4,"type":"Binary","extra":{"Binary":"lico","Ext":"","ID":"lico-build"}}]`,
			path: "/Users/user/Git/lico/dist/lico-build_darwin_arm64/lico",
		},
	}

	tmpfile, err := os.CreateTemp(os.TempDir(), "goreleaser.json")
	if err !=nil{
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	for _, test := range testcase{
		if _, err := tmpfile.Write([]byte(test.json)); err != nil{
			t.Fatal(err)
		}
		
		stdout, _, err := tester.RunCmdWithStdout(artifactCmd, tmpfile.Name())

		if err != nil{
			t.Fatal(err)
		}

		if stdout != test.path{
			t.Errorf("expected %s, but got %s", test.path, stdout)
		}
	}
}

func TestNewcmdCmd(t *testing.T){
	current, err := os.Getwd()
	if err != nil{
		t.Fatal(err)
	}

	basefile := filepath.Clean(filepath.Join(current, "../misc/cmd.go.template"))

	testoutfile, err := os.CreateTemp(os.TempDir(), "testcmd.go")
	if err != nil{
		t.Fatal(err)
	}
	defer os.Remove(testoutfile.Name())

	testcase := []string{
		"testcmd",
	}

	for _, test := range testcase{
		_, _, err :=  tester.RunCmdWithStdout(newcmdCmd, basefile, testoutfile.Name(), test)
		if err != nil{
			t.Fatal(err)
		}
		
		buf := make([]byte, 64)
		_, err = testoutfile.Read(buf)
		if err != nil{
			t.Fatal(err)
		}

		if len(strings.TrimSpace(string(buf))) == 0{
			t.Errorf("expected not empty, but got empty")
		}

	}
}
