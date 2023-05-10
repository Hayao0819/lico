package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Abs(base, path string) (string, error) {
	var err error

	currentDir, err := os.Getwd()
	if err != nil {
		return path, err
	}

	// 起点移動
	if !IsEmpty(base) {
		err = os.Chdir(base)
		if err != nil {
			return path, err
		}
	}

	// 変換
	path = ReplaceTilde(path)
	path, err = filepath.Abs(path)
	if err != nil {
		return path, err
	}

	// ディレクトリを戻る
	err = os.Chdir(currentDir)
	if err != nil {
		return path, err
	}
	return path, nil
}

// ~/を置き換え
func ReplaceTilde(path string) string {
	return strings.Replace(path, "~", GetHomeDir(), 1)
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

func RemoveLine(path string, targetLineNo int)error{
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
		}
	}
	return WriteLines(newFileLine, path)
}

func AppendLine(line , path string)error{
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err !=nil{
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(line))
	return err
}
