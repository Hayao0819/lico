package cmd

import (
	"fmt"
	"os"

	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
)


var repoDir string = "~/.lico/repo"
var listFile string = "~/.lico/lico.list"
var homeDir string
var repoPathBase, homePathBase string

var root *cobra.Command = rootCmd()

func rootCmd() *cobra.Command {
	licoOpt := false

	var cmd = &cobra.Command{
		Use:   "lico",
		Short: "OS非依存なドットファイル管理ツール",
		Long: `licoはOSに依存しないドットファイル管理マネージャーです。
	独自の設定ファイルを用いてホームディレクトリ以下の
	設定ファイルを1つのGitリポジトリで管理します。
	テンプレート記法を用いて柔軟な設定が可能です。`,
		SilenceUsage: true, //コマンド失敗時に使い方を表示しない
		RunE: func(cmd *cobra.Command, args []string) error {
			if licoOpt {
				fmt.Print(lico())
			} else {
				fmt.Fprintln(os.Stderr, "コマンドを指定してください。詳細はlico helpを参照してください。")
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&listFile, "list", "l", listFile, "ファイルリストを指定します")
	cmd.PersistentFlags().StringVarP(&repoDir, "repo", "r", repoDir, "リポジトリディレクトリを指定します")
	cmd.Flags().BoolVarP(&licoOpt, "lico", "", licoOpt, "")
	cmd.Flags().MarkHidden("lico")

	return cmd
}

func common() error {
	// 重要なパスを正規化
	var err error
	listFile, err = utils.Abs("", listFile)
	if err != nil {
		return err
	}
	//fmt.Printf("リスト: %v\n", listFile)

	repoDir, err = utils.Abs("", repoDir)
	if err != nil {
		return err
	}
	//fmt.Printf("リポジトリ: %v\n", repoDir)

	homeDir, err=os.UserHomeDir()
	if err != nil {
		return err
	}

	homePathBase = homeDir
	repoPathBase = repoDir

	return nil
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
