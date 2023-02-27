package cmd

import (
	//"fmt"

	"fmt"
	"os"

	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/conf"
	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
	"github.com/jedib0t/go-pretty/v6/table"
)

type status struct {
	key   string
	value interface{}
}

func loadStatus() ([]status, error) {
	r := []status{}

	// ディレクトリ
	r = append(r, status{key: "ConfigCloned", value: hasCorrectRepoDir()})

	// リポジトリパス
	r = append(r, status{key: "RepoDir", value: *repoDir})

	// リストファイル
	r = append(r, status{key: "ListFile", value: *listFile})

	// リポジトリ
	if repoList, err := getRepoUrl(); err == nil {
		r = append(r, status{key: "RemoteList", value: repoList})
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	// 設定済みリンクの数
	list, err := conf.ReadConf(*listFile)
	if err !=nil{
		return []status{}, err
	}
	var configuredLink, missingLink int
	for _,l:= range *list{
		if e := l.CheckSymLink(); e==nil{
			configuredLink++
		}else{
			missingLink++
		}
	}

	r = append(r, status{key: "FileNum", value: len(*list)})
	r = append(r, status{key: "ConfiguredLink", value: configuredLink})
	r = append(r, status{key: "MissingLink", value: missingLink})

	return r, nil
}

func (s *status) string() string {
	return fmt.Sprintf("%v=%v\n", s.key, s.value)
}

func showTextStatus() error {
	slist , err := loadStatus()
	if err !=nil{
		return err
	}
	for _, s := range slist {
		fmt.Print(s.string())
	}
	return nil
}

func showTableStatus() error {
	/*
		+--------------+-----------------------------------------------------------------------+
		| ConfigCloned | true                                                                  |
		| RepoDir      | /Users/hayao/.lico/repo                                               |
		| ListFile     | /Users/hayao/.lico/lico.list                                          |
		| RemoteList   | [git@github.com:Hayao0819/lico.git git@github.com:Hayao0819/dotfiles] |
		+--------------+-----------------------------------------------------------------------+
	*/

	slist , err := loadStatus()
	if err !=nil{
		return err
	}

	t := table.NewWriter()
	for _, s := range slist {
		t.AppendRow(table.Row{s.key, s.value})
	}

	fmt.Println(t.Render())

	return nil
}

func statusCmd() *cobra.Command {

	var textMode = false

	cmd := cobra.Command{
		Use:   "status",
		Short: "ステータスを表示します",
		Long: `以下の、現在のlicoとシステムの状況を表示します。
・初期化されているかどうか
・同期されているファイル
・同期されていないがリストに存在するファイル
・同期されておらず、システムに別のものが存在するファイル
・現在の環境変数`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if textMode {
				return showTextStatus()
			} else {

				return showTableStatus()
			}
		},
	}

	cmd.Flags().BoolVarP(&textMode, "text", "t", textMode, "テキストモード")

	return &cmd
}

func init() {
	root.AddCommand(statusCmd())
}
