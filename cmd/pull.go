package cmd

import (
	"fmt"
	"os"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
	"github.com/Hayao0819/lico/cmd/common"
)

func pullCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "pull",
		Short:   "設定ファイルリポジトリを最新の状態に更新します",
		Long:    `設定ファイルを管理しているGitリポジトリ内でgit pullを実行します`,
		Args:    cobra.NoArgs,
		Aliases: []string{"init"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.HasCorrectRepoDir() {
				fmt.Fprintln(os.Stderr, "リポジトリがありません。cloneコマンドを用いて初期化してください。")
			} else {
				if err := utils.RunCmd("git", "-C", vars.RepoDir, "pull"); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(pullCmd)
	AddCommand(&cmd)
}
