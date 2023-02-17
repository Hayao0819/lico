package cmd

import (
	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
)

func cloneCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "clone GitURL",
		Short: "設定ファイルリポジトリを取得します",
		Long: `設定ファイルを管理しているGitリポジトリを取得します。
リポジトリは管理ディレクトリ以下にクローンされます。`,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"init"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cloneFrom := args[0]
			gitCmd := utils.MakeCmd("git", "clone", cloneFrom, repoDir)
			err := gitCmd.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	root.AddCommand(cloneCmd())
}
