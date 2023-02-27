package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
)

func cloneCmd() *cobra.Command {
	localPathMode := false

	cmd := cobra.Command{
		Use:   "clone GitURL",
		Short: "設定ファイルリポジトリを取得します",
		Long: `設定ファイルを管理しているGitリポジトリを取得します。
リポジトリは管理ディレクトリ以下にクローンされます。`,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"init"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if hasCorrectRepoDir() {
				return runCmd(pullCmd)
			}

			cloneFrom := args[0]

			if localPathMode {
				return fmt.Errorf("(まだ実装して)ないです。")

				// リポジトリを削除
				// Todo: シンボリックリンクを貼った後にrmrepoを実行すると実体を削除してしまう問題を修正する
				os.RemoveAll(*repoDir)
				if err := os.Symlink(cloneFrom, *repoDir); err != nil {
					return err
				}
			} else if err := utils.RunCmd("git", "clone", cloneFrom, *repoDir); err != nil {
				return err
			}

			if hasCorrectRepoDir() {
				fmt.Println("リポジトリを取得しました。\nsetコマンドを用いて同期を開始してください。")
			} else {
				return errors.New("何らかの理由でリポジトリを初期化できませんでした")
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&localPathMode, "local", "", localPathMode, "ローカルディレクトリをリポジトリとして使う")

	return &cmd
}

func init() {
	root.AddCommand(cloneCmd())
}
