package cmd

import (
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"os"
)

var repoDir *string = &vars.RepoDir
var listFile *string = &vars.ListFile
var homeDir *string = &vars.HomeDir
var repoPathBase *string = &vars.RepoPathBase
var homePathBase *string = &vars.HomePathBase

func common() error {
	// 重要なパスを正規化
	var err error
	vars.ListFile, err = utils.Abs("", vars.ListFile)
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

	homePathBase = homeDir
	repoPathBase = repoDir

	return nil
}
