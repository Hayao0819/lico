package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func rmLinkCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "rmlink [ファイル]",
		Short:   "リンクを削除します",
		Long:    `ホームディレクトリからリンクを削除します。管理ディレクトリ内から実体を削除することはありません。`,
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(os.Stderr, "あとで実装する")
			return nil
		},
	}

	return &cmd
}

func init() {
	root.AddCommand(rmLinkCmd())
}
