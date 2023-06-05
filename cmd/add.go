package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hayao0819/lico/conf"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	var noTemplate bool

	cmd := cobra.Command{
		Use:   "add [flags] homeFile repoFile",
		Short: "ファイルを追加します",
		Long:  `ファイルをlicoの管理対象に追加し、存在しない場合は新たに作成します。`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			entry := conf.NewEntry(p.New(args[1]), p.New(args[0]))

			// Entry一覧を生成
			list, err := conf.ReadConf()
			if err != nil {
				return err
			}

			// Entryに既に登録されていないか確認
			hasHomeFile, _ := list.HasHomeFile(entry.HomePath)   //Todo: 存在しない場合に作成
			hasRepoFile, err := list.HasRepoFile(entry.RepoPath) //Todo: 存在しない場合に作成
			if err != nil {
				return err
			}
			if hasHomeFile || hasRepoFile {
				return errors.New("this file has been managed. Please unregister it first")
			}

			// Ignoreに含まれているか確認
			ignoreList := conf.ReadIgnoreList()
			if status, reg := ignoreList.MatchEntry(entry); status {
				return fmt.Errorf("this path is ignored in \"%s\"", reg)
			}

			// ファイル一覧に追記
			lf, err := os.OpenFile(vars.GetList(), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				return err
			}
			defer lf.Close()

			itemStr, err := entry.String(!noTemplate)
			if err != nil {
				return err
			}

			fmt.Fprint(lf, itemStr)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&noTemplate, "notp", "n", false, "テンプレート構文への置き換えを無効化")

	return &cmd
}

func init() {
	cmd := CmdFunc(addCmd)
	addCommand(&cmd)
}
