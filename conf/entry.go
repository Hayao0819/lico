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
}

// 新しいEntryを生成します
func NewEntry(repoPath, homePath p.Path) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

// Entryを行数付きで作成します
func NewEntryWithIndex(repoPath, homePath p.Path, index int) Entry {
	return Entry{RepoPath: repoPath, HomePath: homePath, Index: index}
}

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

// CreatedListに追記
func addEntryToCreatedList(path p.Path) error {
	if !filepath.IsAbs(path.String()) {
		return errors.New("it should be absolute path")
	}
	return utils.WriteLines([]string{path.String()}, vars.CreatedListFile)
}

// リンクを作成する
func (entry *Entry) MakeSymLink() error {
	// ホームパス
	link, err := entry.HomePath.Abs(*vars.HomePathBase)
	if err != nil {
		return err
	}

	// リポジトリパス
	orig, err := entry.RepoPath.Abs(*vars.RepoPathBase)
	if err != nil {
		return err
	}

	// 確認
	if err := entry.CheckSymLink(); err == nil {
		return nil
		//}else{
		//fmt.Fprintln(os.Stderr, err)
	}

	if !orig.Exists() {
		return vars.ErrNotExist
	}

	if err := os.Symlink(orig.String(), link.String()); err == nil {
		if err := addEntryToCreatedList(link); err != nil {
			os.Remove(link.String())
			return err
		}
		fmt.Printf("%v ==> %v\n", orig.String(), link.String())
		return nil
	} else {
		return err
	}

}

// リンクを削除する
func (entry *Entry) RemoveSymLink() error {
	link, err := entry.FormatRepo()
	if err != nil {
		return err
	}
	if !link.Exists() {
		return vars.ErrNotExist
	}
	if !utils.IsSymlink(link.String()) {
		return vars.ErrNotSymlink
	}

	created, err := ReadCreatedList(vars.CreatedListFile)
	if err != nil {
		return err
	}

	// createdリストに含まれているかどうか
	if res, err := created.HasHomeFile(link); err != nil {
		// createdリストの取得に失敗
		return err
	} else if !res {
		// リストに含まれていない
		return vars.ErrNotManaged
	}

	if  os.Remove(link.String()) != nil{
		// 削除に失敗
		return err
	}

	creatd_entry, err := created.GetItemFromPath(link)
	if err !=nil{
		return err
	}

	// createdlistから該当行を削除
	if utils.RemoveLine(vars.CreatedListFile, creatd_entry.Index) !=nil{
		return err
	}
	return nil
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

