package cmd

import (
	"errors"
	"fmt"
	"io"
	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

func rootCmd(stdin io.Reader, stdout io.Writer, args ...string) *cobra.Command {
	licoOpt := false
	showVersion := false

	cmd := &cobra.Command{
		Use:   "lico",
		Short: "OS非依存なドットファイル管理ツール",
		Long: `licoはOSに依存しないドットファイル管理マネージャーです。
独自の設定ファイルを用いてホームディレクトリ以下の
設定ファイルを1つのGitリポジトリで管理します。
テンプレート記法を用いて柔軟な設定が可能です。`,
		SilenceUsage: true, //コマンド失敗時に使い方を表示しない
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return common.Normalize()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion {
				return common.RunCmd(*GetCommand("version"))
			}

			if licoOpt {
				fmt.Print(common.Lico())
			} else {
				return errors.New("コマンドを指定してください。詳細はlico helpを参照してください。")
			}
			return nil
		},
	}

	// Add commands
	for _, c := range cmdList {
		cmd.AddCommand((*c)())
	}

	// Configure I/O
	cmd.SetIn(stdin)
	cmd.SetOut(stdout)

	// Add flags
	cmd.PersistentFlags().StringVarP(&vars.List, "list", "l", "", "ファイルリストを指定します")
	cmd.PersistentFlags().StringVarP(&vars.RepoDir, "repo", "r", vars.RepoDir, "リポジトリディレクトリを指定します")
	cmd.PersistentFlags().BoolVarP(&showVersion, "version", "", false, "バージョン情報を表示します")
	cmd.PersistentFlags().StringVarP(&vars.Created, "created-list", "", "", "作成されたリンクを保存するファイルを指定します")
	cmd.Flags().MarkHidden("created-list")
	cmd.Flags().BoolVarP(&licoOpt, "lico", "", licoOpt, "")
	cmd.Flags().MarkHidden("lico")

	// help command
	cmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "ヘルプを表示します",
		//Long: ``,
		Args: cobra.NoArgs,
		RunE: func(childcmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	})

	

	return cmd
}

func Execute(stdin io.Reader, stdout io.Writer, args ...string) error {
	err := rootCmd(stdin, stdout, args...).Execute()
	if err != nil {
		return err
	}
	return nil
}
