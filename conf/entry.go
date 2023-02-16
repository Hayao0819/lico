package conf

import (
	//"errors"
	//"fmt"
	"os"

	"github.com/Hayao0819/lico/errs"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
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
func (entry *Entry) MakeSymLink() error {
	link := entry.HomePath
	orig := entry.RepoPath
	if entry.CheckSymLink() == nil {
		return nil
	}

	if !orig.Exists() {
		return errs.ErrNotExist
	}

	err := os.Symlink(orig.String(), link.String())
	if err != nil {
		return err
	}
	return nil
}

// リンクが正常に設定されているかチェックする
func (entry *Entry) CheckSymLink() error {
	link := entry.HomePath.String()
	if !utils.Exists(link) {
		return errs.ErrNotExist
	}

	if !utils.IsSymlink(link) {
		return errs.ErrNotSymlink
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
		return errs.ErrLinkToDiffFile
	}

	return nil
}

// パスがリポジトリファイルに含まれているかどうか
func HasRepoFile(entries *List, path p.Path) (bool, error) {
	for _, entry := range *entries {
		result, err := p.Is(entry.RepoPath, path)
		if err != nil {
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}

// パスがホームファイルに含まれているかどうか
func HasHomeFile(entries *List, path p.Path) (bool, error) {
	for _, entry := range *entries {
		result, err := p.Is(entry.HomePath, path)
		if err != nil {
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}
