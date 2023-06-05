package common

import (
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)

func Normalize() error {
	// 重要なパスを正規化
	err := error(nil)
	vars.RepoDir, err = utils.Abs("", vars.RepoDir)
	if err != nil {
		return err
	}

	// ignoreをコピー
	ignore := vars.Ignore
	vars.Ignore=[]string{}
	for _, i := range ignore{
		fixed, err := utils.Abs("", i)
		if err == nil{
			vars.Ignore = append(vars.Ignore, fixed)
		}
	}


	// いろんなパスを正規化
	for _, v := range []*string{&vars.List, &vars.Created, &vars.PkgList,} {
		if utils.IsEmpty(*v) {
			continue
		}

		*v, err = utils.Abs(vars.RepoDir, *v)
		if err != nil{
			return err
		}
	}
	return nil
}
