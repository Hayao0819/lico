package dotfile

type Entry struct{
	RepoPath string
	HomePath string
}


func NewEntry(repoPath, homePath string)(Entry){
	return Entry{RepoPath: repoPath, HomePath: homePath}
}

