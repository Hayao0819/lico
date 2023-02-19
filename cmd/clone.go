package cmd

import (
	"errors"
	"fmt"

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
			if hasCorrectRepoDir() {
				gitCmd := utils.MakeCmd("git", "-C", repoDir, "pull")
				if err := gitCmd.Run(); err != nil {
					return err
				}
			} else {
				cloneFrom := args[0]
				gitCmd := utils.MakeCmd("git", "clone", cloneFrom, repoDir)
				if err := gitCmd.Run(); err != nil {
					return err
				}
			}

			if hasCorrectRepoDir() {
				fmt.Println("リポジトリを取得しました。\nsetコマンドを用いて同期を開始してください。")
			} else {
				return errors.New("何らかの理由でリポジトリを初期化できませんでした")
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	root.AddCommand(cloneCmd())
}
