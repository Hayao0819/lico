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


func NewEntry(repoPath, homePath Path)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

func (entry *Entry) String ()(string){
	return fmt.Sprintf("%v:%v\n", entry.RepoPath, entry.HomePath)
}

func HasRepoFile(entries *[]Entry, path Path)(bool, error){
	fullPath, err := path.Abs()
	if err != nil{
		return false, err
	}

	// フルパスで一致したなら一致
	for _,entry := range *entries{
		repoPathFile, err := entry.RepoPath.Abs()
		if err != nil{
			continue
		}
		if repoPathFile == fullPath {
			return true, nil
		}
	}
	return false, nil
}

func HasHomeFile(entries *[]Entry, path Path)(bool, error){
	fullPath, err := path.Abs()
	if err != nil{
		return false, err
	}

	// フルパスで一致したなら一致
	for _,entry := range *entries{
		homePathFile, err := entry.HomePath.Abs()
		if err != nil{
			continue
		}
		if homePathFile == fullPath {
			return true, nil
		}
	}
	return false, nil
}
