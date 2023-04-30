package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

type CmdFunc (func() *cobra.Command)

var cmdList = map[string](*CmdFunc){}

// Add a command to the command list.
func addCommand(cmd *CmdFunc) {
	cmdList[(*cmd)().Name()] = cmd
}

// Get a command from the command list.
func GetSubCmd(name string) *CmdFunc {
	return cmdList[name]
}

// コマンド内から別のコマンドを実行する
func RunSubCmdFromCmd(name string, cmd *cobra.Command, args ...string) error {
	subcmdfunc := GetSubCmd(name)
	if subcmdfunc == nil {
		return nil
	}
	subcmd := (*subcmdfunc)()
	subcmd.SetArgs(args)
	subcmd.SetOut(cmd.OutOrStdout())
	return subcmd.Execute()
}

// 出力を指定してコマンドを実行する
func RunSubCmdWithIO(name string, stdout, stderr io.Writer, args ...string) error {
	cmdfunc := GetSubCmd(name)
	if cmdfunc == nil {
		return nil
	}
	cmd := (*cmdfunc)()

	// argsのフラグをcmdに渡す
	cmd.SetArgs(args)

	// 出力を指定
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	return cmd.Execute()
}
