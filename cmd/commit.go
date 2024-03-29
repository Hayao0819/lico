package cmd

import (
	"fmt"
	"os"

	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
)

func commitCmd() *cobra.Command {
	var gitflags string
	cmd := cobra.Command{
		Use:   "commit",
		Short: "変更をコミットします",
		Long:  `設定ファイルの変更をコミットします`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.HasCorrectRepoDir() {
				fmt.Fprintln(os.Stderr, "リポジトリがありません。cloneコマンドを用いて初期化してください。")
			}

			var err error
			gitArgs := []string{"-C", vars.RepoDir}
			if !utils.IsEmpty(gitflags) {
				gitArgs = append(gitArgs, gitflags)
			}
			if err = utils.RunCmd("git", append(gitArgs, "add", "-A")...); err != nil {
				return err
			}
			if len(args) == 0 {
				if err = utils.RunCmd("git", append(gitArgs, "commit")...); err != nil {
					return err
				}
			} else {
				if err = utils.RunCmd("git", append(gitArgs, "commit", "-m", args[0])...); err != nil {
					return err
				}
			}
			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(commitCmd)
	addCommand(&cmd)
}
