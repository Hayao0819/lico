/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"errors"
	"fmt"

	"github.com/Hayao0819/lico/conf"
	df "github.com/Hayao0819/lico/dotfile"
	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
)


var delLineMode bool = false

func unregisterCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "unregister",
		Short: "ファイルを管理対象から除外",
		Long: `指定されたファイルをlicoの管理対象から除外します。
		ファイルは削除されず、関連付けのみ解除されます。
		デフォルトでは設定ファイルの該当箇所をコメントアウトします。
		行を完全に削除するには-dを用いてください。`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error{
			// Entry一覧を生成
			list, err := conf.ReadConf(listFile)
			if err != nil{
				return err
			}

			targetPath := df.Path(args[0])

			targetItem := list.GetItemFromPath(targetPath)
			if targetItem == nil{
				return fmt.Errorf("no such file is registered")
			}
			//fmt.Println(targetItem.Index)
			err = utils.CommentOut(listFile, targetItem.Index)
			return err
		},
	}

	return &cmd
}

func init() {
	cmd := unregisterCmd()
	rootCmd.AddCommand(cmd)
	cmd.Flags().BoolVarP(&delLineMode , "del-line", "d", false, "Help message for toggle")
}
