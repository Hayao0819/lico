package cmd

import (
	"errors"
	"fmt"

	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

var root *cobra.Command = rootCmd()

func rootCmd() *cobra.Command {
	licoOpt := false
	showVersion := false

	var cmd = &cobra.Command{
		Use:   "lico",
		Short: "OS非依存なドットファイル管理ツール",
		Long: `licoはOSに依存しないドットファイル管理マネージャーです。
独自の設定ファイルを用いてホームディレクトリ以下の
設定ファイルを1つのGitリポジトリで管理します。
テンプレート記法を用いて柔軟な設定が可能です。`,
		SilenceUsage: true, //コマンド失敗時に使い方を表示しない
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion {
				return runCmd(versionCmd)
			}

			if licoOpt {
				fmt.Print(lico())
			} else {
				return errors.New("コマンドを指定してください。詳細はlico helpを参照してください。")
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&vars.List, "list", "l", vars.GetList(), "ファイルリストを指定します")
	cmd.PersistentFlags().StringVarP(&vars.RepoDir, "repo", "r", vars.RepoDir, "リポジトリディレクトリを指定します")
	cmd.PersistentFlags().BoolVarP(&showVersion, "version", "", false, "バージョン情報を表示します")
	cmd.PersistentFlags().StringVarP(&vars.Created, "created-list", "", vars.GetCreated(), "作成されたリンクを保存するファイルを指定します")
	cmd.Flags().MarkHidden("created-list")
	cmd.Flags().BoolVarP(&licoOpt, "lico", "", licoOpt, "")
	cmd.Flags().MarkHidden("lico")

	return cmd
}

// コマンドを実行します
// 引数: バージョン, コミット, 日付
func Execute() error {
	var err error
	err = common()
	if err != nil {
		return err
	}

	err = root.Execute()
	if err != nil {
		return err
	}
	return nil
}
