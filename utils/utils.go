package utils

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func IsDir(path string)(bool){
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}


func IsFile(path string)(bool){
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}


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
