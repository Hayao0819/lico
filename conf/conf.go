package conf

import (
	"bufio"
	"errors"
	"strings"

	"os"

	df "github.com/Hayao0819/lico/dotfile"
)


func ReadConf(path string)(*[]df.Entry, error){
	file, err := os.Open(path)
	if err != nil{
		return nil, errors.New("cannot open file")
	}

	var entrySlice []df.Entry
	var entry df.Entry
	var splited []string
	var repoPath string
	var homePath string

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		splited = strings.Split(scanner.Text(), ":")
		repoPath = splited[0]
		homePath = splited[1]

		entry = df.NewEntry(repoPath, homePath)
		entrySlice = append(entrySlice, entry)
	}
	return &entrySlice,nil
}
