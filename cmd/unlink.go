package cmd

import (
	//"errors"
	"fmt"

	"github.com/Hayao0819/lico/conf"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/spf13/cobra"
)

func unlinkCmd() *cobra.Command {

	var delLineMode bool = false
	var noEditListMode bool = false

	cmd := cobra.Command{
		Use:   "unlink",
		Short: "ファイルを管理対象から除外",
		Long: `指定されたファイルをlicoの管理対象から除外します。
		ファイルは削除されず、関連付けのみ解除されます。
		デフォルトでは設定ファイルの該当箇所をコメントアウトします。
		行を完全に削除するには-dを用いてください。`,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"unregister"},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Entry一覧を生成
			list, err := conf.ReadConf(listFile)
			if err != nil {
				return err
			}

			targetPath := p.NewAbs(args[0])

			targetItem := list.GetItemFromPath(targetPath)
			if targetItem == nil {
				return fmt.Errorf("no such file is registered")
			}

			if !noEditListMode {
				if err = utils.CommentOut(listFile, targetItem.Index); err != nil {
					return err
				}
			}

			if err = rmLinkCmd().RunE(rmLinkCmd(), []string{targetPath.String()}); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&delLineMode, "delline", "d", delLineMode, "コメントアウトの代わりに行を削除します")
	cmd.Flags().BoolVarP(&noEditListMode, "noedit", "", noEditListMode, "リストファイルを編集しません")

	return &cmd
}

func init() {
	cmd := unlinkCmd()
	root.AddCommand(cmd)
}
