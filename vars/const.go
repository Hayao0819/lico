package vars

import (
	"github.com/Hayao0819/lico/utils"
)

var RepoPathBase string
var HomePathBase string

func Init(repodir string)error{
	var err error

	HomePathBase, err = func()(string, error){
		osinfo, err := utils.GetOSEnv()
		if err !=nil{
			return "", err
		}
		return osinfo["Home"], nil
	}()
	if err !=nil{
		return err
	}

	RepoPathBase = repodir

	return nil
}
