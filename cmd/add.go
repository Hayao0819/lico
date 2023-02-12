/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"
	"errors"
	"os"

	"github.com/Hayao0819/lico/conf"
	df "github.com/Hayao0819/lico/dotfile"
	"github.com/spf13/cobra"
)



func addCmd ()(*cobra.Command){
	cmd := cobra.Command{
		Use:   "add [flags] repoFile hooeFile",
		Short: "ファイルを追加します",
		Long: `ファイルをlicoの管理対象に追加し、存在しない場合は新たに作成します。`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			entry := df.NewEntry(df.Path(args[0]) , df.Path(args[1]))

			// Entry一覧を生成
			list, err := conf.ReadConf(listFile)
			if err != nil{
				return err
			}

			// Entryに既に登録されていないか確認
			hasHomeFile, _ := df.HasHomeFile(list.GetEntries(), entry.HomePath) //Todo: 存在しない場合に作成
			hasRepoFile, err := df.HasRepoFile(list.GetEntries(), entry.RepoPath) //Todo: 存在しない場合に作成
			if err !=nil { return err; }
			//fmt.Printf("hasHomeFile=%v hasRepoFile=%v\n", hasHomeFile, hasRepoFile)
			if hasHomeFile || hasRepoFile {
				return errors.New("this file has been managed. Please unregister it first")
			}

			// ファイル一覧に追記
			lf, err := os.OpenFile(listFile, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				return err
			}
			defer lf.Close()
			//fmt.Fprint(lf, entry.String())

			return nil
		},
	}

	return &cmd
}

func init() {
	rootCmd.AddCommand(addCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
