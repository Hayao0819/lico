package conf

import (
	//"errors"
	//"fmt"
	"fmt"
	"os"

	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)

// 設定ファイルの1行に対応した構造体
type Entry struct {
	RepoPath p.Path
	HomePath p.Path
	Index    int
}

// 新しいEntryを生成します
func NewEntry(repoPath, homePath p.Path) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

func NewEntryWithIndex(repoPath, homePath p.Path, index int) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath, Index: index}
}

// repoPathが存在するかどうかを確認する
func (entry *Entry) ExistsRepoPath() bool {
	_, err := os.Stat(string(entry.RepoPath))
	return err == nil
}

// リンクを作成する
func (entry *Entry) MakeSymLink(homeBasePath string, repoBasePath string) error {
	// ホームパス
	var link p.Path
	if linkF, err := entry.HomePath.Abs(homeBasePath); err != nil {
		return err
	} else if link, err = Format(linkF.String()); err != nil {
		return err
	}

	// リポジトリパス
	var orig p.Path
	if origF, err := entry.RepoPath.Abs(repoBasePath); err != nil {
		return err
	} else if orig, err = Format(origF.String()); err != nil {
		return err
	}

	// 確認
	if entry.CheckSymLink() == nil {
		return nil
	}

	if !orig.Exists() {
		return vars.ErrNotExist
	}

	if err := os.Symlink(orig.String(), link.String()); err == nil {
		fmt.Printf("%v ==> %v\n", orig.String(), link.String())
		return nil
	} else {
		return err
	}

}

// リンクが正常に設定されているかチェックする
func (entry *Entry) CheckSymLink() error {
	link := entry.HomePath.String()
	if !utils.Exists(link) {
		return vars.ErrNotExist
	}

	if !utils.IsSymlink(link) {
		return vars.ErrNotSymlink
	}

	readlink, err := os.Readlink(link)
	if err != nil {
		return err
	}

	isSameFile, err := p.Is(p.New(readlink), entry.RepoPath)
	if err != nil {
		return err
	}

	if !isSameFile {
		return vars.ErrLinkToDiffFile
	}

	return nil
}
