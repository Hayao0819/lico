package cmd

import (
	//"fmt"

	"fmt"
	"os"

	//"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
)

func showTextStatus() error {
	// ディレクトリ
	fmt.Printf("ConfigCloned=%v\n", hasCorrectRepoDir())

	// リポジトリパス
	fmt.Printf("RepoDir=%v\n", *repoDir)

	// リストファイル
	fmt.Printf("ListFile=%v\n", *listFile)

	// リポジトリ
	if repoList , err := getRepoUrl(); err ==nil{
		fmt.Printf("RemoteList=%v\n", repoList)
	}else{
		fmt.Fprintln(os.Stderr, err)
	}

	// 
	
	return nil
}

func statusCmd() *cobra.Command {

	var textMode = true

	cmd := cobra.Command{
		Use:   "status",
		Short: "ステータスを表示します",
		Long: `以下の、現在のlicoとシステムの状況を表示します。
・初期化されているかどうか
・同期されているファイル
・同期されていないがリストに存在するファイル
・同期されておらず、システムに別のものが存在するファイル
・現在の環境変数`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if textMode{
				return showTextStatus()
			}else{
				return nil
			}
		},
	}

	cmd.Flags().BoolVarP(&textMode, "text", "t", textMode,"テキストモード")

	return &cmd
}

func init() {
	root.AddCommand(statusCmd())
}
