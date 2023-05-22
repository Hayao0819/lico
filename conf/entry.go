package conf

import (
	//"errors"
	//"fmt"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)

// 設定ファイルの1行に対応した構造体
type Entry struct {
	RepoPath p.Path
	HomePath p.Path
	Index    int //0からスタートする行数
	Option   *EntryOption
}

// 新しいEntryを生成します
func NewEntry(repoPath, homePath p.Path) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

// Entryを行数付きで作成します
func NewEntryWithIndex(repoPath, homePath p.Path, index int) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath, Index: index}
}

// Entryを行数、オプション付きで作成します
/*
func NewFullEntry(repoPath, homePath p.Path, index int, opt *EntryOption)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath, Index: index, Option: opt}
}
*/

// repoPathが存在するかどうかを確認する
func (entry *Entry) ExistsRepoPath() bool {
	_, err := os.Stat(string(entry.RepoPath))
	return err == nil
}

func (entry *Entry) FormatHome() (p.Path, error) {
	return entry.HomePath.Abs(vars.HomeDir)
}

func (entry *Entry) FormatRepo() (p.Path, error) {
	return entry.RepoPath.Abs(*vars.RepoPathBase)
}

func (item *Entry) String(replace bool) (string, error) {
	var (
		repo, home p.Path
	)
	var err error

	if replace {
		repo, err = replaceToTemplate(item.RepoPath.String())
		if err != nil {
			return "", err
		}

		home, err = replaceToTemplate(item.HomePath.String())
		if err != nil {
			return "", err
		}
	} else {
		repo = item.RepoPath
		home = item.HomePath
	}
	return fmt.Sprintf("%v:%v\n", repo, home), nil
}

// CreatedListに追記
func addEntryToCreatedList(path p.Path) error {
	if !filepath.IsAbs(path.String()) {
		return errors.New("it should be absolute path")
	}
	return utils.AppendLine(path.String(), vars.GetCreated())
}
