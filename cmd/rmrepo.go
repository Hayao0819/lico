package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
	"github.com/manifoldco/promptui"
)

func rmrepoCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "rmrepo",
		Short: "リポジトリを削除します",
		Long: `リポジトリをディレクトリごと削除します。
現時点のバージョンで、コミットされていない内容や変更されたファイルの確認は行わないため十分に注意してください。
リンクの削除も行わないので、注意してください。`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			noConfirm := false
			
			if ! noConfirm{
				p := promptui.Select{
					Label: fmt.Sprintf("%vを削除します。よろしいですか？", *repoDir),
					Items: []string{"Yes", "No"},
				}
				_, selected, err := p.Run()
				if err !=nil{
					return err
				}

				if selected!="Yes"{
					return nil
				}
			}

			if err := os.RemoveAll(*repoDir); err!=nil{
				return err
			}

			fmt.Fprintln(os.Stderr, "リポジトリを削除しました。cloneコマンドを用いて初期化し直してください。")

			return nil
		},
	}

	return &cmd
}

func init() {
	root.AddCommand(rmrepoCmd())
}
