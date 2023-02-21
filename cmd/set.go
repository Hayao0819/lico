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
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := conf.ReadConf(*listFile)
			var errlist []error
			if err != nil {
				return err
			}

			for _, entry := range *list {
				if entry.CheckSymLink() == nil {
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
