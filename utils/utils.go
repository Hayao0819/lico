package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	//"strings"
	"sort"
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

// ファイルの内容を読みとって文字列配列を返します
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// 指定したファイルの指定した行の行頭に#を追加します
func CommentOut(path string, targetLineNo int) error {
	var newFileLine []string
	fileLines, err := ReadLines(path)
	if err != nil {
		return err
	}
	lineNo := 0
	for _, line := range fileLines {
		lineNo++
		if targetLineNo != lineNo {
			newFileLine = append(newFileLine, line)
			continue
		} else {
			newFileLine = append(newFileLine, fmt.Sprintf("#%s", line))
		}
	}

	return WriteLines(newFileLine, path)
}

// コマンドの存在確認
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// 文字列配列を長さでソート
func SortWithLen(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j])
	})
	return arr
}
