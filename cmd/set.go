package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/utils"
	"errors"

	"github.com/spf13/cobra"
)

func setCmd() *cobra.Command {
	dry_run := false

	cmd := cobra.Command{
		Use:   "set",
		Short: "シンボリックリンクを作成",
		Long: `リストに従ってシンボリックリンクを作成します
もし不正なファイルが設定されていた場合、そのファイルは無視して続行されます。
`,
		RunE: func(cmd *cobra.Command, ars []string) error {
			

			// get conf
			list, err := conf.ReadConf()
			var errlist []error
			if err != nil {
				return err
			}

			for _, entry := range *list {
				// dry-runが有効な場合は、実際にリンクを作成せずに、作成する予定のリンクを表示する
				show_msg := dry_run

				if err := entry.CheckSymLink(); err == nil {
					continue
				}

				if ! dry_run {
					err := entry.MakeSymLink()
					if err != nil {
						errlist = append(errlist, err)
					}else{
						show_msg = true
					}
				}
				
				if show_msg {
					cmd.Printf("%v ==> %v\n", entry.RepoPath.String(), entry.HomePath.String())
				}
			}

			if len(errlist) == 0 {
				return nil
			} else {
				for _, err := range errlist {
					fmt.Fprintln(os.Stderr, err)
				}
				return errors.New(strings.Join(func(errlist []error) []string {
					var rtn []string
					for _, err := range errlist {
						rtn = append(rtn, err.Error())
					}
					return rtn
				}(errlist), "\n"))
			}

		},
	}

	cmd.Flags().BoolVarP(&dry_run, "dry-run", "d", false, "dry run")

	return &cmd
}

func init() {
	cmd := CmdFunc(setCmd)
	addCommand(&cmd)
}
