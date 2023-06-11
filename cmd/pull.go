package cmd

import (
	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
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
				cmd.PrintErrln("リポジトリがありません。cloneコマンドを用いて初期化してください。")
			} else {
				if err := utils.RunCmd("git", "-C", vars.RepoDir, "pull"); err != nil {
					return err
				}
			}

			if getStatus("MissingLink") != 0{
				cmd.PrintErrln("新しいリンクがあります。setコマンドを用いて設定してください。")
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(pullCmd)
	addCommand(&cmd)
}
