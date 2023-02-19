package utils

import (
	"os"
	"os/exec"
	//"strings"
	"sort"
)

// 文字列配列を長さでソート
func SortWithLen(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j])
	})
	return arr
}

/*
func ForEach(arr []interface{}, runFunc func(int, interface{})([]interface{}))([]interface{}){
	rtn := []interface{}{}
	for index, item := range arr{
		rtn = append(rtn, runFunc(index, item))
	}
	return rtn
}

func ForEachStop(arr []interface{}, runFunc func(int, interface{})(error))(error){
	for index, item := range arr{
		err := runFunc(index, item)
		if err == nil{
			return err
		}
	}
	return nil
}
*/

func MakeCmd(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd
}
