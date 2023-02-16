package path

import (
	"os"
	"path/filepath"

	"github.com/Hayao0819/lico/utils"
)

// いくつかのメソッドを持ったファイルパス
type Path string

// 　ファイルのStatを返します
func (path *Path) Stat() (os.FileInfo, error) {
	return os.Stat(string(*path))

}

// ファイルの絶対パスを返します
func (path *Path) Abs() (Path, error) {
	str, err := filepath.Abs(string(*path))
	return New(str), err
}

// パスを文字列に変換
func (path *Path) String() string {
	return string(*path)
}

// utils.Existsを用いてファイルが存在するかどうかを確認します
func (path *Path) Exists() bool {
	return utils.Exists(path.String())
}
