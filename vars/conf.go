package vars

import (
	"path"
	"os"
	"github.com/Hayao0819/lico/utils"
)


var RepoDir string = "~/.lico/repo"

/*
var BaseListFile string = "~/.lico/repo/lico.list"
var IgnoreListFile string = "~/.lico/repo/lico.ignore"
var CreatedListFile string = "~/.lico/created.list"
var PkgListFile string = "~/.lico/repo/lico-pkgs-2.json"
*/

var (
	List string = ""
	Ignore string = ""
	Created string = ""
	PkgList string = ""
	HomeDir string = ""
)

var (
	RepoPathBase *string = &RepoDir
	HomePathBase *string = &HomeDir
)

func GetList()string{
	if ! utils.IsEmpty(List){
		return List
		
	}
	return path.Join(RepoDir + "/lico.list")
}

func GetIgnore()string{
	if ! utils.IsEmpty(Ignore){
		return Ignore
		
	}
	return path.Join(RepoDir + "/lico.ignore")
}

func GetCreated()string{
	if ! utils.IsEmpty(Created){
		return Created
		
	}
	return path.Join(RepoDir + "/created.list")
}

func GetPkgList()string{
	if ! utils.IsEmpty(PkgList){
		return PkgList
		
	}
	return path.Join(RepoDir + "/lico-pkgs-2.json")
}

func init(){
	HomeDir,_ = os.UserHomeDir()
}

