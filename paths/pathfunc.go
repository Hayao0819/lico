package path

import (
	"path/filepath"
)

func New(pathS string) Path {
	return Path(pathS)
}

func NewAbs(pathS string) Path {
	absPath, _ := filepath.Abs(pathS)
	return Path(absPath)
}

// 2つのパスが共通のファイルを指しているかどうかを確認します
func Is(path1, path2 Path) (bool, error) {
	path1Abs, err := path1.Abs("")
	if err != nil {
		return false, err
	}
	path2Abs, err := path2.Abs("")
	if err != nil {
		return false, err
	}

	if path1Abs == path2Abs {
		return true, nil
	} else {
		return false, nil
	}
}
