package dotfile

import (
	//"errors"
	"fmt"
	"os"
)

type Entry struct{
	RepoPath Path
	HomePath Path
}


type Path string

func (path *Path) Stat()(os.FileInfo, error){
	return os.Stat(string(*path))
	
}


func NewEntry(repoPath, homePath Path)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

func (entry *Entry) String ()(string){
	return fmt.Sprintf("%v:%v\n", entry.RepoPath, entry.HomePath)
}

func HasRepoFile(entries *[]Entry, path Path)(bool, error){
	pathInfo, err := path.Stat()
	if err != nil{
		return false, err
	}

	// os.Statで同じオブジェクトを返却したら一致
	// もしエラーの場合でもフルパスで一致したなら一致
	for _,entry := range *entries{
		repoPathFile, err := entry.RepoPath.Stat()
		if err != nil{
			if string(path) == string(entry.RepoPath){
				return true, nil
			}
			continue
		}

		if pathInfo == repoPathFile {
			return true, nil
		}
	}
	return false, nil
}

func HasHomeFile(entries *[]Entry, path Path)(bool, error){
	pathInfo, err := path.Stat()
	if err != nil{
		return false, err
	}

	// os.Statで同じオブジェクトを返却したら一致
	// もしエラーの場合でもフルパスで一致したなら一致
	for _,entry := range *entries{
		homePathFile, err := entry.HomePath.Stat()
		if err != nil{
			if string(path) == string(entry.HomePath){
				return true, nil
			}
			continue
		}

		if pathInfo == homePathFile {
			return true, nil
		}
	}
	return false, nil
}
