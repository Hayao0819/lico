package cmd

import (
	"fmt"
	"os"

	"github.com/Hayao0819/lico/conf"
	"github.com/spf13/cobra"

	"golang.org/x/exp/slices"
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

	cmd.AddCommand(oldlinkCmd())
	cmd.AddCommand(allCmd())
	cmd.AddCommand(ignoreCmd())
	cmd.AddCommand(duplicateCmd())

	return &cmd
}

func allCmd() *cobra.Command {
	rm_all := false
	verbose_msg := false

	cmd := cobra.Command{
		Use:   "all",
		Short: "全ての問題を修正",
		Long:  `全ての問題を自動で修正します。`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			oldlink_args := []string{}

			if rm_all {
				oldlink_args = append(oldlink_args, "--all")
			}

			if verbose_msg {
				oldlink_args = append(oldlink_args, "--verbose")
			}

			// oldlink
			oldlink_cmd := oldlinkCmd()
			oldlink_cmd.SetArgs(oldlink_args)
			err := oldlink_cmd.Execute()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&rm_all, "all", "a", false, "全ての問題を修正します(ファイルを削除します)")
	cmd.Flags().BoolVarP(&verbose_msg, "verbose", "v", false, "詳細なメッセージを表示します")

	return &cmd
}

func oldlinkCmd() *cobra.Command {
	rm_all := false
	rm_nolink := false
	rm_broken := false
	rm_unregistered := false
	verbose_msg := false

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
				if verbose_msg {
					println("Checking: " + e.HomePath.String())
				}

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
	cmd.Flags().BoolVarP(&verbose_msg, "verbose", "v", false, "詳細なメッセージを表示します")

	return &cmd
}

// Todo: 重複したファイルの自動修正
func duplicateCmd() *cobra.Command{
	cmd := cobra.Command{
		Use: "duplicate",
		Aliases: []string{"dup"},
		Short: "重複したファイルを確認",
		Long: `リストに登録されているが重複しているファイルを表示・除外します。`,
		Args: cobra.NoArgs,
	}

	createdCmd := cobra.Command{
		Use: "created",
		Short: "重複した作成済みリンクの一覧を確認",
		Long: `リストに登録されているが重複している作成済みリンクを表示・除外します。`,

		RunE: func(cmd *cobra.Command, args []string) error {
			created, err := conf.ReadCreatedList()
			if err != nil {
				return err
			}

			checked := []string{}
			for _, e := range *created {
				if slices.Contains(checked, e.HomePath.String()) {
					fmt.Fprintf(os.Stderr, "重複した作成済みリンク: %s\n", e.HomePath)
				} else {
					checked = append(checked, e.HomePath.String())
				}
			}

			return nil
		},
	}

	cmd.AddCommand(&createdCmd)

	return &cmd
}

func ignoreCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ignore",
		Short: "本来無視されるべきファイルを確認",
		Long:  `リストに登録されているが本来無視されるべきファイルを表示・除外します。`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := conf.ReadConf()
			if err != nil {
				return err
			}

			ignore, err := conf.ReadIgnoreList()
			if err != nil {
				return err
			}

			for _, e := range *list {
				if r, _ := ignore.MatchEntry(e); r {
					cmd.Println(e.HomePath.String())
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
