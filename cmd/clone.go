package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/vars"
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
			if common.HasCorrectRepoDir() {
				return common.RunCmd(pullCmd)
			}

			cloneFrom := args[0]

			if localPathMode {
				//return fmt.Errorf("(まだ実装して)ないです。")

				// リポジトリを削除
				os.RemoveAll(vars.RepoDir)
				if err := os.Symlink(cloneFrom, vars.RepoDir); err != nil {
					return err
				}
			} else if err := utils.RunCmd("git", "clone", cloneFrom, vars.RepoDir); err != nil {
				return err
			}

			if common.HasCorrectRepoDir() {
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
	cmd := CmdFunc(cloneCmd)
	AddCommand(&cmd)
}
