package cmd

import (
	//"fmt"

	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"

	//"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/vars"
)

func gitCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "git [git args]...",
		Short: "gitを実行する",
		Long:  `設定ファイルリポジトリ内でGitコマンドを実行します`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.HasCorrectRepoDir() {
				return vars.ErrNoRepoDir
			}

			return utils.RunCmd("git", append([]string{"-C", vars.RepoDir}, args...)...)
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(gitCmd)
	addCommand(&cmd)
}
