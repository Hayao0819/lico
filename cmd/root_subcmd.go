package cmd

import "github.com/spf13/cobra"

type CmdFunc (func() *cobra.Command)

var cmdList = map[string](*CmdFunc){}

func AddCommand(cmd *CmdFunc) {
	cmdList[(*cmd)().Name()] = cmd
}

func GetCommand(name string) *CmdFunc {
	return cmdList[name]
}
