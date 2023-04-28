package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

var editorBin string

func editCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "edit",
		Short: fmt.Sprintf("%vを手動で編集", vars.GetList()),
		Long: `リストファイルを手動で編集します。

エディタはオプションもしくは環境変数"EDITOR"で指定されたものが起動されます。
もしエディタを認識できない場合はViを起動します。`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if editorBin == "" {
				return errors.New("cannot find an editor")
			}

			editorRun := exec.Command("sh", "-c", "--", fmt.Sprintf("%v %v", editorBin, vars.GetList()))
			editorRun.Stdout = os.Stdout
			editorRun.Stdin = os.Stdin
			editorRun.Stderr = os.Stderr
			err := editorRun.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&editorBin, "editor", "e",
		func() string {
			env := strings.TrimSpace(os.Getenv("EDITOR"))
			if len(env) == 0 {
				return "vi"
			} else {
				return env
			}
		}(),
		"編集に利用するエディタを設定します")

	return &cmd
}

func init() {
	cmd := CmdFunc(editCmd)
	AddCommand(&cmd)
}
