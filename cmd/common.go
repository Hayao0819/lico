package cmd

import (
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"os"
)

var repoDir *string = &vars.RepoDir
var listFile *string = &vars.BaseListFile

//var createdListFile *string = &vars.CreatedListFile
//var homeDir *string = &vars.HomeDir
//var repoPathBase *string = &vars.RepoPathBase
//var homePathBase *string = &vars.HomePathBase

func common() error {
	// 重要なパスを正規化
	var err error
	vars.BaseListFile, err = utils.Abs("", vars.BaseListFile)
	if err != nil {
		return err
	}
	//fmt.Printf("リスト: %v\n", listFile)

	vars.RepoDir, err = utils.Abs("", vars.RepoDir)
	if err != nil {
		return err
	}
	//fmt.Printf("リポジトリ: %v\n", repoDir)

	vars.HomeDir, err = os.UserHomeDir()
	if err != nil {
		return err
	}

	vars.CreatedListFile, err = utils.Abs("", vars.CreatedListFile)
	if err != nil {
		return err
	}

	//homePathBase = homeDir
	//repoPathBase = repoDir

	return nil
}
