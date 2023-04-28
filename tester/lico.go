package tester

import (
	//"os/exec"
	"os"
	"github.com/Hayao0819/lico/cmd"
)


func RunLico (args ...string)error{
	args = append([]string{os.Args[0]}, args...)  // 引数の先頭にファイル名を追加
	return cmd.Execute(os.Stdin, os.Stdout, args...)
}

func MakeSymLinkInExample() error {
	return RunLico("set")
}


