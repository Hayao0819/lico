package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/conf"
	p "github.com/Hayao0819/lico/paths"
	"github.com/spf13/cobra"
)

func rmFileCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "rmfile [ファイル]",
		Short: "リポジトリからファイルを削除します",
		Long: `リポジトリからファイルを削除します。リンクも一緒に削除されます。
ファイルパスの指定にはリンクと実ファイルの両方を用いることができます。`,
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get List
			var list *conf.List
			var err error
			if list, err = conf.ReadConf(); err != nil {
				return err
			}

			// エラーリスト
			var errList []error

			for _, arg := range args {
				targetPath := p.New(arg)
				targetEntry, err := list.GetItemFromPath(targetPath)
				if err != nil {
					errList = append(errList, err)
				} else {
					//abs, err := targetEntry.HomePath.Abs(vars.HomePathBase)
					//abs, err := formatRepoPath(&targetEntry.RepoPath)
					abs, err := targetEntry.FormatHome()
					if err != nil {
						errList = append(errList, err)
					}
					if err := os.Remove(abs.String()); err != nil {
						errList = append(errList, err)
					}
				}
			}

			// あとでちゃんと実装する
			//rmLinkCmd().RunE(rmFileCmd(), args)
			common.RunCmd(rmLinkCmd, args...)

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

	return &cmd
}

func init() {
	cmd := CmdFunc(rmFileCmd)
	AddCommand(&cmd)
}
