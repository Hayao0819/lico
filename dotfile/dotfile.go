package dotfile

import (
	//"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/errs"
)

// 設定ファイルの1行に対応した構造体
type Entry struct{
	RepoPath Path
	HomePath Path
}

// いくつかのメソッドを持ったファイルパス
type Path string

func (path *Path) Stat()(os.FileInfo, error){
	return os.Stat(string(*path))
	
}

// ファイルの絶対パスを返します
func (path *Path)Abs()(string, error){
	return filepath.Abs(string(*path))
}

func (path *Path)String()(string){
	return string(*path)
}

func NewPath(pathS string)Path{
	return Path(pathS)
}

func NewAbsPath(pathS string)Path{
	absPath, _ := filepath.Abs(pathS)
	return Path(absPath)
}

// 2つのパスが共通のファイルを指しているかどうかを確認します
func PathIs(path1 , path2 Path)(bool, error){
	path1Abs, err := path1.Abs()
	if err != nil{
		return false, err
	}
	path2Abs , err := path2.Abs()
	if err != nil{
		return false, err
	}

	if path1Abs == path2Abs{
		return true,nil
	}else{
		return false, nil
	}
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
	// Todo 実行する
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
