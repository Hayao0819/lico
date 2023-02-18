package utils

import (
	"os"
	"os/exec"
	"strings"
)


// 文字列が正常なディレクトリへのパスかどうかを確認します
func IsDir(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

// 文字列が正常なファイルパスかどうかを調べます
func IsFile(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}

// シンボリックリンクかどうか
// 参考: https://github.com/eihigh/filetest
// Thanks eihigh <eihigh.contact@gmail.com>
func IsSymlink(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode()&os.ModeSymlink != 0
}

// ファイルが存在するかどうか
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// コマンドの存在確認
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}
