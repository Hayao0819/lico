package dotfile

import (
	//"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Entry struct{
	RepoPath Path
	HomePath Path
}


type Path string

func (path *Path) Stat()(os.FileInfo, error){
	return os.Stat(string(*path))
	
}

func (path *Path)Abs()(string, error){
	return filepath.Abs(string(*path))
}

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

func NewEntry(repoPath, homePath Path)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

func (entry *Entry) String ()(string){
	return fmt.Sprintf("%v:%v\n", entry.RepoPath, entry.HomePath)
}

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
