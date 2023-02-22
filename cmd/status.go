package cmd

import (
	//"fmt"

	"fmt"

	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
)

func statusCmd() *cobra.Command {

	showInitilizedStatus := func()error{
		repoDirStatus := hasCorrectRepoDir()
		if repoDirStatus{
			fmt.Println("リポジトリ同期済み")
			url , err := getRepoUrl()
			if err !=nil{
				return err
			}
			fmt.Println(url)
			
		}
		return nil
	}

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
			stdin, stderr, err := utils.RunCmdAndGet("sh", "-c", "echo My-stdout; echo My-stderr 1>&2")
			fmt.Println(stdin[0])
			fmt.Println(stderr[0])
			fmt.Println(err)
			
			return showInitilizedStatus()
			//return nil
		},
	}

	return &cmd
}

func init() {
	root.AddCommand(statusCmd())
}
