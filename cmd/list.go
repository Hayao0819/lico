package cmd

import (
	"errors"
	"fmt"
	"os"

	//"strings"

	//"os"

	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// listCmd represents the list command
func listCmd() *cobra.Command {
	absPathMode := false
	relPathMode := false

	showIndexNo := false

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

				textToPrint := ""

				parsedRepoPath, err := formatRepoPath(&entry.RepoPath)
				if err != nil {
					return err
				}
				parsedHomePath, err := formatHomePath(&entry.HomePath)
				if err != nil {
					return err
				}

				if absPathMode {
					//fmt.Printf("%v%s%v\n", parsedRepoPath, listSeparator, parsedHomePath)
					textToPrint=parsedRepoPath.String() + listSeparator + parsedHomePath.String()
					//continue
				}else if relPathMode {
					parsedRelRepoPath, err := parsedRepoPath.Rel(*repoPathBase)
					if err != nil {
						return err
					}

					parsedRelHomePath, err := parsedHomePath.Rel(*homePathBase)
					if err != nil {
						return err
					}

					textToPrint=parsedRelRepoPath.String() + listSeparator+parsedRelHomePath.String()
					//fmt.Printf("%v%s%v\n", parsedRelRepoPath, listSeparator, parsedRelHomePath)
					//continue
				}


				if utils.IsEmpty(textToPrint){
					fmt.Fprintln(os.Stderr, "このメッセージが出力されることはありえないはずです。バグを作者に報告してください。")
					return errors.New("no mode specified")
				}


				if showIndexNo{
					textToPrint = fmt.Sprintf("%v: %v", entry.Index+1, textToPrint) //entry.Indexは0からスタートします
				}

				fmt.Println(textToPrint)
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&absPathMode, "abs", "a", absPathMode, "テンプレートを解釈してフルパスで表示")
	cmd.Flags().BoolVarP(&relPathMode, "rel", "s", relPathMode, "テンプレートを解釈して相対パスで表示(デフォルト)")
	cmd.Flags().StringVarP(&listSeparator, "sep", "", listSeparator, "リストの区切り文字を指定")
	cmd.Flags().BoolVarP(&nullSeparator, "null", "", false, "リストの区切り文字にヌル文字を指定")
	cmd.Flags().BoolVarP(&showIndexNo, "lineno","", showIndexNo, "設定されている行数を表示" )

	// Allow --line-no
	cmd.Flags().SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		if name == "line-no"{
			name="lineno"
		}
		return pflag.NormalizedName(name)
	})

	return &cmd
}

func init() {
	root.AddCommand(listCmd())
}
