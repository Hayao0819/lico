package path

import (
	"os"

	"github.com/Hayao0819/lico/utils"
	"path/filepath"
)

// いくつかのメソッドを持ったファイルパス
type Path string

// 　ファイルのStatを返します
func (path *Path) Stat() (os.FileInfo, error) {
	return os.Stat(string(*path))

}

// ファイルの絶対パスを返します
func (path *Path) Abs(base string) (Path, error) {
	str, err := utils.Abs(base, string(*path))
	return New(str), err
}

// ファイルの相対パスを返します
func (path *Path) Rel(base string) (Path, error) {
	rel, err := filepath.Rel(base, path.String())
	return New(rel), err
}

// パスを文字列に変換
func (path *Path) String() string {
	return string(*path)
}

// utils.Existsを用いてファイルが存在するかどうかを確認します
func (path *Path) Exists() bool {
	return utils.Exists(path.String())
}

func (path *Path) IsSymlink() bool {
	return utils.IsSymlink(path.String())
}
