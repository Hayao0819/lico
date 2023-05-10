package common

import (
	"errors"
	"os"

	"github.com/Hayao0819/lico/vars"
)

func GlobalMode()error{
	vars.GlobalMode = true

	// root権限で実行しているか確認
	if os.Getuid() != 0 {
		return errors.New("global modeはroot権限で実行してください。")
	}

	// ディレクトリの設定
	vars.RepoDir = "/etc/lico/repo"
	vars.Created = "/etc/lico/created.list"

	return nil
}
