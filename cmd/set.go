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

	cmd := cobra.Command{
		Use:   "set",
		Short: "シンボリックリンクを作成",
		Long: `リストに従ってシンボリックリンクを作成します
もし不正なファイルが設定されていた場合、そのファイルは無視して続行されます。
`,
		RunE: func(cmd *cobra.Command, ars []string) error {
			run_by_admin := false
			// Check root
			if os.Getegid() == 0{
				run_by_admin = true
			}

			// get conf
			list, err := conf.ReadConf()
			var errlist []error
			if err != nil {
				return err
			}

			for _, entry := range *list {
				
				if run_by_admin && ! entry.Option.WithRoot {
					// rootとして実行かつWithRootではない
					if err := entry.CheckSymLink(); err !=nil{
						fmt.Fprintln(os.Stderr, "一般ユーザーとして実行して全てのシンボリックリンクを作成してください")
					}
					continue
				}else if ! run_by_admin && entry.Option.WithRoot{
					// 一般ユーザーとして実行かつWithRoot
					if err := entry.CheckSymLink(); err !=nil{
						fmt.Fprintln(os.Stderr, "Rootとして実行して全てのシンボリックリンクを作成してください")
					}
					continue
				}

				err := entry.MakeSymLink()
				if err != nil {
					errlist = append(errlist, err)
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

	return &cmd
}

func init() {
	root.AddCommand(setCmd())
}
