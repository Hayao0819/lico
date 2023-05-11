package tester

import (
	//"os/exec"
	"bytes"
	"os"
	"strings"

	"github.com/Hayao0819/lico/cmd"
	"github.com/spf13/cobra"
)


func RunLico (args ...string)error{
	args = append([]string{os.Args[0]}, args...)  // 引数の先頭にファイル名を追加
	return cmd.Execute(os.Stdin, os.Stdout, args...)
}

func MakeSymLinkInExample() error {
	return RunLico("set")
}


func RunCmdWithStdout(f func()(cmd *cobra.Command), args ...string)(string, string, error){
	cmd := f()

	// set stdout, stderr
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	// deal with args
	if len(args) == 0{
		cmd.SetArgs([]string{"--"})
	}else{
		cmd.SetArgs(args)
	}

	// run
	err := cmd.Execute()
	//return stdout.String(), stderr.String(), err

	return strings.TrimSuffix(stdout.String(), "\n"), strings.TrimSuffix(stderr.String(), "\n"), err
}
