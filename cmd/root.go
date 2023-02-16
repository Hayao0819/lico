package cmd

import (
	"fmt"
	"os"

	"encoding/base64"
	"github.com/spf13/cobra"
)

var listFile string
var root *cobra.Command = rootCmd()


func rootCmd ()(*cobra.Command){
	shamiko := false

	var cmd = &cobra.Command{
		Use:   "lico",
		Short: "OS非依存なドットファイル管理ツール",
		Long: `licoはOSに依存しないドットファイル管理マネージャーです。
	独自の設定ファイルを用いてホームディレクトリ以下の
	設定ファイルを1つのGitリポジトリで管理します。
	テンプレート記法を用いて柔軟な設定が可能です。`,
		SilenceUsage: true, //コマンド失敗時に使い方を表示しない
		Run: func(cmd *cobra.Command, args []string){
			if shamiko{
				dec, _ := base64.StdEncoding.DecodeString("44KI44GP6KaL44Gf44KJ44GT44Gu5a2Q44Gq44KT44GL5aSJ44Gq44KT55Sf44GI44Go44KL77yB77yBCg==")
				fmt.Print(string(dec))
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&listFile, "list", "l", "~/.lico/lico.list", "ファイルリストを指定します")
	cmd.PersistentFlags().BoolVarP(&shamiko, "shamiko", "", shamiko, "")
	cmd.PersistentFlags().MarkHidden("shamiko")

	return cmd
}



// コマンドを実行します
func Execute() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
