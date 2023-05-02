package cmd

import (
	"errors"
	"fmt"
	"os"

	//"strings"

	//"os"

	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"

	//"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// listCmd represents the list command
func listCmd() *cobra.Command {
	// abs or rel
	absPathMode := false
	relPathMode := false

	// show index number
	showIndexNo := false

	// separator
	listSeparator := " ==> "
	nullSeparator := false

	// show created list
	showCreatedList := false

	cmd := cobra.Command{
		Use:   "list",
		Short: "ドットファイルの一覧を表示",
		Long:  ``,
		PreRunE: func(cmd *cobra.Command, args []string) error {
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

			if nullSeparator {
				listSeparator = string([]byte{0})
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if showCreatedList{
				createdList, err := conf.ReadCreatedList()
				if err != nil {
					return err
				}
				for _, entry := range *createdList {
					parsed, err := entry.FormatRepo()
					if err != nil {
						return nil
					}
					if absPathMode {
						fmt.Println(parsed.String())
						continue
					}else{
						rel, err := parsed.Rel(*vars.HomePathBase)
						if err !=nil{
							return err
						}
						fmt.Println(rel.String())
					}
				}
				return nil
			}else{
				list, err := conf.ReadConf()
				if err != nil {
					return err
				}
				for _, entry := range *list {
					textToPrint := ""

					//parsedRepoPath, err := formatRepoPath(&entry.RepoPath)
					parsedRepoPath, err := entry.FormatRepo()
					if err != nil {
						return err
					}
					//parsedHomePath, err := formatHomePath(&entry.HomePath)
					parsedHomePath, err := entry.FormatHome()
					if err != nil {
						return err
					}

					if absPathMode {
						//cmd.Printf("%v%s%v\n", parsedRepoPath, listSeparator, parsedHomePath)
						textToPrint = parsedRepoPath.String() + listSeparator + parsedHomePath.String()
						//continue
					} else if relPathMode {
						parsedRelRepoPath, err := parsedRepoPath.Rel(*vars.RepoPathBase)
						if err != nil {
							return err
						}

						parsedRelHomePath, err := parsedHomePath.Rel(*vars.HomePathBase)
						if err != nil {
							return err
						}

						textToPrint = parsedRelRepoPath.String() + listSeparator + parsedRelHomePath.String()
						//cmd.Printf("%v%s%v\n", parsedRelRepoPath, listSeparator, parsedRelHomePath)
						//continue
					}

					if utils.IsEmpty(textToPrint) {
						fmt.Fprintln(os.Stderr, "このメッセージが出力されることはありえないはずです。バグを作者に報告してください。")
						return errors.New("no mode specified")
					}

					if showIndexNo {
						textToPrint = fmt.Sprintf("%v: %v", entry.Index+1, textToPrint) //entry.Indexは0からスタートします
					}

					cmd.Println(textToPrint)
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&showCreatedList, "created", "c", showCreatedList, "作成されたリンクの一覧を表示")
	cmd.Flags().BoolVarP(&absPathMode, "abs", "a", absPathMode, "テンプレートを解釈してフルパスで表示")
	cmd.Flags().BoolVarP(&relPathMode, "rel", "s", relPathMode, "テンプレートを解釈して相対パスで表示(デフォルト)")
	cmd.Flags().StringVarP(&listSeparator, "sep", "", listSeparator, "リストの区切り文字を指定")
	cmd.Flags().BoolVarP(&nullSeparator, "null", "", false, "リストの区切り文字にヌル文字を指定")
	cmd.Flags().BoolVarP(&showIndexNo, "lineno", "", showIndexNo, "設定されている行数を表示")

	// Allow --line-no
	cmd.Flags().SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		if name == "line-no" {
			name = "lineno"
		}
		return pflag.NormalizedName(name)
	})

	return &cmd
}

func init() {
	cmd := CmdFunc(listCmd)
	addCommand(&cmd)
}
