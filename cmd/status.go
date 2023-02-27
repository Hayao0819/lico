package cmd

import (
	//"fmt"

	"fmt"
	"os"

	//"github.com/Hayao0819/lico/utils"
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

func loadStatus() []status {
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

	return r
}

func (s *status) string() string {
	return fmt.Sprintf("%v=%v\n", s.key, s.value)
}

func showTextStatus() error {
	for _, s := range loadStatus() {
		fmt.Print(s.string())
	}
	return nil
}

func showTableStatus() error {
	/*
		----------------------------
		|   ConfigCloned   |  true  |
		|   RepoDir        |  hoge  |
		-----------------------------

	*/

	t := table.NewWriter()
	for _, s := range loadStatus() {
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
