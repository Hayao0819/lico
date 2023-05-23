package cmd

import (
	"errors"
	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func printEntryWithAbs(cmd *cobra.Command, entry *conf.Entry, sep string, showline bool)error{
	entry, err := entry.Format()
	if err != nil {
		return err
	}
	cmd.Println(entry.RepoPath.String() + sep + entry.HomePath.String())
	return nil
}

func printEntryWithRel(cmd *cobra.Command, entry *conf.Entry, sep string, showline bool)error{
	entry, err := entry.Format()
	if err != nil {
		return err
	}
	relHome, err := entry.HomePath.Rel(*vars.HomePathBase)
	if err != nil {
		return err
	}
	relRepo, err := entry.RepoPath.Rel(*vars.RepoPathBase)
	if err != nil {
		return err
	}
	cmd.Println(relRepo.String() + sep + relHome.String())
	return nil
}

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
			if showCreatedList {
				list, err := conf.ReadConf()
				if err != nil {
					return err
				}
				createdList, err := conf.ReadCreatedList()
				if err != nil {
					return err
				}
				for _, entry := range *createdList {
					parsed, err := entry.FormatHome()
					if err != nil {
						return nil
					}
					hashome, _ := list.HasHomeFile(parsed)
					e, err := list.GetItemFromPath(parsed)
					if err != nil || e == nil {
						hashome=false
					}
					if hashome{
						// リストに登録されているシンボリックリンク
						if absPathMode {
							printEntryWithAbs(cmd, e, listSeparator, showIndexNo)
						} else {
							printEntryWithRel(cmd, e, listSeparator, showIndexNo)
						}
					}else{
						// リストに登録されていないシンボリックリンク
						if absPathMode {
							cmd.Println(parsed.String())
							continue
						} else {
							rel, err := parsed.Rel(*vars.HomePathBase)
							if err != nil {
								return err
							}
							cmd.Println(rel.String())
						}
					}
				}
				return nil
			} else {
				list, err := conf.ReadConf()
				if err != nil {
					return err
				}
				for _, entry := range *list {
					if absPathMode {
						printEntryWithAbs(cmd, &entry, listSeparator, showIndexNo)
					} else if relPathMode {
						printEntryWithRel(cmd, &entry, listSeparator, showIndexNo)
					}
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
