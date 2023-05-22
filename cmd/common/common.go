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

	for _, v := range []*string{&vars.List, &vars.Ignore, &vars.Created, &vars.PkgList} {
		if utils.IsEmpty(*v) {
			continue
		}

		*v, err = utils.Abs(vars.RepoDir, *v)
		if err != nil {
			*v = utils.ReplaceTilde(*v)
		}
	}
	return nil
}
