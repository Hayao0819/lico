package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/Hayao0819/lico/conf"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

func rmLinkCmd() *cobra.Command {
	rmAll := false

	cmd := cobra.Command{
		Use:   "rmlink [ファイル]",
		Short: "リンクを削除します",
		Long:  `ホームディレクトリからリンクを削除します。管理ディレクトリ内から実体を削除することはありません。`,
		Args: func(cmd *cobra.Command, args []string) error {
			if rmAll {
				if err := cobra.NoArgs(cmd, args); err != nil {
					return err
				}
			} else {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Aliases: []string{},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get List
			var list *conf.List
			var err error
			var rmList []string
			var errList []error
			if list, err = conf.ReadConf(*listFile); err != nil {
				return err
			}

			if rmAll {
				// rmAllで警告を出したい
				for _, i := range *list {
					rmList = append(rmList, i.HomePath.String())
				}
			} else {
				rmList = args
			}

			for _, arg := range rmList {
				targetPath := p.New(arg)
				targetEntry := list.GetItemFromPath(targetPath)
				if targetEntry == nil {
					errList = append(errList, vars.ErrNoSuchEntry(arg))
				} else {
					//abs, err := formatHomePath(&targetEntry.HomePath)
					abs, err := targetEntry.FormatHome()
					if err != nil {
						errList = append(errList, err)
					}
					if err := os.Remove(abs.String()); err != nil {
						errList = append(errList, err)
					}
				}
			}

			// 最終処理
			if len(errList) == 0 {
				return nil
			} else {
				errStringList := func() []string {
					var rtn []string
					for _, err := range errList {
						rtn = append(rtn, err.Error())
					}
					return rtn
				}()
				return errors.New(strings.Join(errStringList, "\n"))
			}
		},
	}

	cmd.Flags().BoolVarP(&rmAll, "all", "", rmAll, "全てのリンクを削除します")

	return &cmd
}

func init() {
	root.AddCommand(rmLinkCmd())
}
