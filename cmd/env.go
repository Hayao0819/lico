/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Hayao0819/lico/utils"
)



func envCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "env [Name]",
		Short: "リストで利用できる変数を表示",
		Long: `リストで利用できる変数を表示します。
リストはGolangのテンプレート構文をサポートしており、環境依存の値を埋め込むことが可能です。
このコマンドではテンプレート構文内で利用可能な変数の一覧を表示します。
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string)error {
			// envを取得
			env, err := utils.GetOSEnv()
			if err !=nil{
				return err
			}


			if len(args)==0{
				for index, value := range env{
					fmt.Printf("%v = %v\n", index, value)
				}
			}else if len(args)==1{
				for index, value := range env{
					if index == args[0]{
						fmt.Printf("%v\n", value)
					}
				}
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	rootCmd.AddCommand(envCmd())
}
