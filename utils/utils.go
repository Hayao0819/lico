package utils

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

// 文字列が正常なディレクトリへのパスかどうかを確認します
func IsDir(path string)(bool){
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

// 文字列が正常なファイルパスかどうかを調べます
func IsFile(path string)(bool){
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
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
func CommentOut(path string, targetLineNo int)error{
	var newFileLine []string
	fileLines, err := ReadLines(path)
	if err!=nil{
		return err
	}
	lineNo:=0
	for _, line := range fileLines {
		lineNo++
		if targetLineNo != lineNo{
			newFileLine = append(newFileLine, line)
			continue
		}else{
			newFileLine = append(newFileLine, fmt.Sprintf("#%s", line))
		}
	}

	return WriteLines(newFileLine, path)
}

//func Write
