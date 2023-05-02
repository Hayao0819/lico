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
		Long:  `同期の問題など、様々なものの異常を表示します。修正するにはオプションを用いて明示的に指示する必要があります。`,
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(oldlinkcmd())

	return &cmd
}

func oldlinkcmd() *cobra.Command {
	rm_all := false
	rm_nolink := false
	rm_broken := false
	rm_unregistered := false

	cmd := cobra.Command{
		Use:   "oldlink",
		Short: "リストと設定されているリンクを同期",
		Long:  "ファイルリストとシステムに設定されているリンクを確認し、古いリンクや壊れているリンクを修正します",
		PreRun: func(cmd *cobra.Command, args []string) {
			if rm_all {
				rm_nolink = true
				rm_broken = true
				rm_unregistered = true
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			created := vars.GetCreated()
			list, err := conf.ReadConf()
			if err != nil {
				return err
			}

			created_list, err := conf.ReadCreatedList()
			if err != nil {
				return err
			}

			//remove_path := []*p.Path{}
			for _, e := range *created_list {
				// Prepare 
				home := e.HomePath
				println("Checking: " + e.HomePath.String())

				// Check if the link is symlink
				if !utils.IsSymlink(home.String()) {
					fmt.Fprintf(os.Stderr, "リンクではない: %s\n", home)
					if rm_nolink {
						if utils.RemoveLine(created, e.Index) != nil {
							return err
						}
					}
					continue
				}

				// Check if the link is broken
				realpathstr, err := os.Readlink(home.String())
				if err != nil || utils.IsEmpty(realpathstr) {
					// 破損したリンク
					fmt.Fprintf(os.Stderr, "破損したリンク: %v\n", home)
					if rm_broken {
						if e.RemoveSymLink() != nil {
							cmd.PrintErr(err)
						}
					}
					continue
				}

				// Check if the link is registered
				if _, err = list.GetItemFromPath(home); err != nil {
					// すでに登録解除されたリンク
					fmt.Fprintf(os.Stderr, "リストにないリンク: %v\n", home)
					
					if rm_unregistered {
						if e.RemoveSymLink() != nil {
							return err
						}
						if utils.RemoveLine(created, e.Index) != nil {
							return err
						}
					}
					continue
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&rm_all, "all", "a", false, "全てのリンクを削除します")
	cmd.Flags().BoolVarP(&rm_nolink, "nolink", "n", false, "リンクではないファイルをリストから除外します")
	cmd.Flags().BoolVarP(&rm_broken, "broken", "b", false, "壊れたリンクを削除します")
	cmd.Flags().BoolVarP(&rm_unregistered, "unregistered", "u", false, "リストにないリンクを削除します")

	return &cmd
}

func init() {
	cmd := CmdFunc(fixCmd)
	addCommand(&cmd)
}
