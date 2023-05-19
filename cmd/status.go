package cmd

import (
	//"fmt"

	"errors"
	"fmt"
	"os"

	"github.com/Hayao0819/lico/cmd/common"
	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"

	//"github.com/Hayao0819/lico/vars"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type status struct {
	key   string
	value interface{}
	desc  string
}

func newStatus(key string, value interface{}, desc string) status {
	return status{key: key, value: value, desc: desc}
}

func loadStatus() []status {
	r := []status{}
	errs := []error{}

	// ディレクトリ
	//r = append(r, status{key: "ConfigCloned", value: hasCorrectRepoDir()})
	r = append(r, newStatus("ConfigCloned", common.HasCorrectRepoDir(), "Dotfilesリポジトリが配置済みかどうか"))

	// リポジトリパス
	//r = append(r, status{key: "RepoDir", value: *repoDir})
	r = append(r, newStatus("RepoDir", vars.RepoDir, "Dotfilesリポジトリのパス"))

	//リポジトリパスがシンボリックリンクかどうか
	//r= append(r, status{key: "IsSymlink", value: utils.IsSymlink(*repoDir)})
	r = append(r, newStatus("IsSymlink", utils.IsSymlink(vars.RepoDir), "RepoDirがシンボリックリンクかどうか"))

	// リストファイル
	//r = append(r, status{key: "ListFile", value: *listFile})
	r = append(r, newStatus("ListFile", vars.GetList(), "リストファイルのパス"))

	// リポジトリ
	if repoList, err := common.GetRepoUrl(); err == nil {
		//r = append(r, status{key: "RemoteList", value: repoList})
		r = append(r, newStatus("RemoteList", repoList, "Dotfilesを管理しているリモートリポジトリ"))
	} else {
		//fmt.Fprintln(os.Stderr, err)
		//return []status{}, err
		errs = append(errs, err)
	}

	// 設定済みリンクの数
	list, err := conf.ReadConf()
	if err != nil {
		//return []status{}, err
		errs = append(errs, err)
	}
	var configuredLink, missingLink int
	if list != nil {
		for _, l := range *list {
			if e := l.CheckSymLink(); e == nil {
				configuredLink++
			} else {
				missingLink++
			}
		}
	} else {
		list = &conf.List{}
	}

	// licoによって作成されたリンクの数
	managedLink := 0
	created, err := conf.ReadCreatedList()
	if err != nil {
		errs = append(errs, err)
	} else {
		managedLink = len(*created)
	}

	//r = append(r, status{key: "FileNum", value: len(*list)})
	//r = append(r, status{key: "ConfiguredLink", value: configuredLink})
	//r = append(r, status{key: "MissingLink", value: missingLink})
	r = append(r,
		newStatus("FileNum", len(*list), "登録されてるリンクの数(OSによって変化する場合があります)"),
		newStatus("ConfiguredLink", configuredLink, "適切に配置されているリンクの数"),
		newStatus("MissingLink", missingLink, "まだ設定されていないリンクの数"),
		newStatus("ManagedLink", managedLink, "licoによって作成されたリンクの数"),
	)

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, e)
		}
		return r
	}
	return r
}

func (s *status) string() string {
	return fmt.Sprintf("%v=%v\n", s.key, s.value)
}

func showTextStatus(cmd *cobra.Command) error {
	slist := loadStatus()
	for _, s := range slist {
		cmd.Print(s.string())
	}
	return nil
}

func showTableStatus(cmd *cobra.Command) error {
	/*
		+--------------+-----------------------------------------------------------------------+
		| ConfigCloned | true                                                                  |
		| RepoDir      | /Users/hayao/.lico/repo                                               |
		| ListFile     | /Users/hayao/.lico/lico.list                                          |
		| RemoteList   | [git@github.com:Hayao0819/lico.git git@github.com:Hayao0819/dotfiles] |
		+--------------+-----------------------------------------------------------------------+
	*/

	slist := loadStatus()

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Key", "Desc", "Value"})
	for _, s := range slist {
		t.AppendRow(table.Row{s.key, s.desc, s.value})
	}
	t.SetStyle(table.StyleBold)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Key", AlignHeader: text.AlignCenter},
		{Name: "Desc", AlignHeader: text.AlignCenter},
		{Name: "Value", AlignHeader: text.AlignCenter},
	})

	cmd.Println(t.Render())

	return nil
}

func showValue(key string, cmd *cobra.Command) error {
	slist := loadStatus()

	for _, s := range slist {
		if s.key == key {
			cmd.Println(s.value)
		}
	}

	return nil
}

func statusCmd() *cobra.Command {

	textMode := false
	tableMode := false

	cmd := cobra.Command{
		Use:   "status",
		Short: "ステータスを表示します",
		Long: `以下の、現在のlicoとシステムの状況を表示します。
・初期化されているかどうか
・同期されているファイル
・同期されていないがリストに存在するファイル
・同期されておらず、システムに別のものが存在するファイル
・現在の環境変数`,
		Args: cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if textMode && tableMode {
				return errors.New("textモードとtableモードは同時に指定できません")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				return showValue(args[0], cmd)
			}

			if textMode {
				return showTextStatus(cmd)
			} else if tableMode {
				return showTableStatus(cmd)
			} else {
				return showTextStatus(cmd)
			}
		},
	}

	cmd.Flags().BoolVarP(&textMode, "text", "t", textMode, "テキストモード")
	cmd.Flags().BoolVarP(&tableMode, "table", "T", tableMode, "テーブルモード")

	return &cmd
}

func init() {
	cmd := CmdFunc(statusCmd)
	addCommand(&cmd)
}
