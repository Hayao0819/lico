package dotfile

import (
	//"errors"
	"fmt"
	"os"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/errs"
)

// 設定ファイルの1行に対応した構造体
type Entry struct{
	RepoPath Path
	HomePath Path
}


// 新しいEntryを生成します
func NewEntry(repoPath, homePath Path)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

// Entryから設定ファイル用文字列を生成します
func (entry *Entry) String ()(string){
	return fmt.Sprintf("%v:%v\n", entry.RepoPath, entry.HomePath)
}


// repoPathが存在するかどうかを確認する
func (entry *Entry) ExistsRepoPath() (bool){
	_, err := os.Stat(string(entry.RepoPath))
	return err == nil
}

// リンクを作成する
func (entry *Entry) MakeSymLink()(error){
	link := entry.HomePath
	orig := entry.RepoPath
	if entry.CheckSymLink() == nil{
		return nil
	}

	if ! orig.Exists(){
		return errs.ErrNotExist
	}

	err := os.Symlink(orig.String(), link.String())
	if err !=nil{
		return err
	}
	return nil
}

// リンクが正常に設定されているかチェックする
func (entry *Entry) CheckSymLink()(error){
	link := entry.HomePath.String()
	if ! utils.Exists(link){
		return errs.ErrNotExist
	}

	if ! utils.IsSymlink(link){
		return errs.ErrNotSymlink
	}

	readlink, err := os.Readlink(link)
	if err != nil{
		return err
	}

	isSameFile, err := PathIs(NewPath(readlink), entry.RepoPath)
	if err !=nil{
		return err
	}

	if ! isSameFile{
		return errs.ErrLinkToDiffFile
	}

	return nil
}

// パスがリポジトリファイルに含まれているかどうか
func HasRepoFile(entries *[]Entry, path Path)(bool, error){
	for _,entry := range *entries{
		result, err := PathIs(entry.RepoPath, path)
		if err != nil{
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}

// パスがホームファイルに含まれているかどうか
func HasHomeFile(entries *[]Entry, path Path)(bool, error){
	for _,entry := range *entries{
		result, err := PathIs(entry.HomePath, path)
		if err != nil{
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}