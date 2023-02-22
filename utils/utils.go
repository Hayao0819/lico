package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

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


func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func RunCmdAndGet(name string, args ...string) ([]string, []string ,error){
	cmd := exec.Command(name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	
	si, sr := func (s, r bytes.Buffer)([]string, []string){
		rtn := [2][]string{}
		for i, b := range []bytes.Buffer{s, r}{
			rtn[i]=func(b bytes.Buffer)([]string){
				return strings.Split(b.String(), "\n") 
			}(b)
		}
		
		return rtn[0], rtn[1]
	}(stdout, stderr)

	return si, sr, err
}

/*
func MergeMap(m ...map[string]interface{}) map[string]interface{} {
    ans := make(map[string]interface{}, 0)

    for _, c := range m {
        for k, v := range c {
            ans[k] = v
        }
    }
    return ans
}
*/

func GetHomeDir() string {
	homedir, _ := os.UserHomeDir()
	return homedir
}
