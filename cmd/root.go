package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var listFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lico",
	Short: "OS非依存なドットファイル管理ツール",
	Long: `licoはOSに依存しないドットファイル管理マネージャーです。
独自の設定ファイルを用いてホームディレクトリ以下の
設定ファイルを1つのGitリポジトリで管理します。
テンプレート記法を用いて柔軟な設定が可能です。`,
	SilenceUsage: true, //コマンド失敗時に使い方を表示しない
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// コマンドを実行します
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lico.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&listFile, "list", "l", "~/.lico/lico.list", "ファイルリストを指定します")
	//rootCmd.PersistentFlags().BoolVarP()
}
