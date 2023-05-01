package cmd

import (
	"fmt"
	"os"

	"github.com/Hayao0819/lico/conf"
	"github.com/spf13/cobra"

	//p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)

func fixCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "fix",
		Short: "様々な問題を修正",
		Long:  `同期の問題など、様々なものを修正し正常に動作させます`,
		Args:  cobra.NoArgs,
		/*
			RunE: func(cmd *cobra.Command, args []string) error {

				return nil
			},
		*/
	}

	cmd.AddCommand(oldlinkcmd())

	return &cmd
}

func oldlinkcmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "oldlink",
		Short: "リストと設定されているリンクを同期",
		Long:  "ファイルリストとシステムに設定されているリンクを確認し、古いリンクや壊れているリンクを修正します",
		RunE: func(cmd *cobra.Command, args []string) error {

			created := vars.GetCreated()

			list, err := conf.ReadConf()
			if err != nil {
				return err
			}

			creatd, err := conf.ReadCreatedList()
			if err != nil {
				return err
			}

			/*
			for _, e := range *list {
				// リンクを作成
				if e.CheckSymLink() != nil {
					if err := e.MakeSymLink(); err != nil {
						return err
					}
				}
			}
			*/

			//remove_path := []*p.Path{}
			for _, e := range *creatd {
				home := e.HomePath
				realpathstr, err := os.Readlink(home.String())

				if !utils.IsSymlink(home.String()) {
					fmt.Fprintf(os.Stderr, "リンクではない: %s\n", home)
					if utils.RemoveLine(created, e.Index) != nil {
						return err
					}
					continue
				}

				if err != nil || utils.IsEmpty(realpathstr) {
					// 破損したリンク
					fmt.Fprintf(os.Stderr, "破損したリンク: %v\n", home)
					//remove_path=append(remove_path, &home)
					continue
				}

				if _, err = list.GetItemFromPath(home); err != nil {
					// すでに登録解除されたリンク
					fmt.Fprintf(os.Stderr, "リストにないリンク: %v\n", home)
					if e.RemoveSymLink() != nil {
						return err
					}
					if utils.RemoveLine(created, e.Index) != nil {
						return err
					}
					continue
				}
			}
			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(fixCmd)
	addCommand(&cmd)
}
