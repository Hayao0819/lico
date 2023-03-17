package cmd

import (
	"errors"
	"fmt"
	"os"
	//"strings"

	//"os"

	"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
func listCmd() *cobra.Command {
	absPathMode := false
	relPathMode := false

	listSeparator := " ==> "
	nullSeparator := false

	cmd := cobra.Command{
		Use:   "list",
		Short: "ドットファイルの一覧を表示",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 引数チェック
			trueN := func(b ...bool) int {
				c := 0
				for _, v := range b {
					if v {
						c++
					}
				}
				return c
			}(absPathMode, relPathMode)
			if trueN == 0 {
				// デフォルト動作
				relPathMode = true
			} else if trueN >= 2 {
				return errors.New("multiple output methods specified")
			}

			if nullSeparator{
				listSeparator=string([]byte{0})
			}


			// 設定ファイルを読み込み
			list, err := conf.ReadConf(*listFile)
			if err != nil {
				//fmt.Fprintln(os.Stderr, err)
				return err
			}

			//listSeparator=strings.ReplaceAll(listSeparator, `\n`, "\n")

			for _, entry := range *list {
				parsedRepoPath, err := formatRepoPath(&entry.RepoPath)
				if err != nil {
					return err
				}
				parsedHomePath, err := formatHomePath(&entry.HomePath)
				if err != nil {
					return err
				}

				if absPathMode {
					fmt.Printf("%v%s%v\n", parsedRepoPath, listSeparator, parsedHomePath)
					continue
				}

				if relPathMode {
					parsedRelRepoPath, err := parsedRepoPath.Rel(*repoPathBase)
					if err != nil {
						return err
					}

					parsedRelHomePath, err := parsedHomePath.Rel(*homePathBase)
					if err != nil {
						return err
					}
					fmt.Printf("%v%s%v\n", parsedRelRepoPath, listSeparator, parsedRelHomePath)
					continue
				}

				fmt.Fprintln(os.Stderr, "このメッセージが出力されることはありえないはずです。バグを作者に報告してください。")
				return errors.New("no mode specified")
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&absPathMode, "abs", "a", absPathMode, "テンプレートを解釈してフルパスで表示")
	cmd.Flags().BoolVarP(&relPathMode, "rel", "s", relPathMode, "テンプレートを解釈して相対パスで表示(デフォルト)")
	cmd.Flags().StringVarP(&listSeparator, "sep", "", listSeparator, "リストの区切り文字を指定")
	cmd.Flags().BoolVarP(&nullSeparator, "null", "", false, "リストの区切り文字にヌル文字を指定")

	return &cmd
}

func init() {
	root.AddCommand(listCmd())
}
