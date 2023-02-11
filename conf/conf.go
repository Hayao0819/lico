package conf

import (
	"bufio"
	"errors"
	"strings"
	"regexp"
	"os"

	df "github.com/Hayao0819/lico/dotfile"
)


func ReadConf(path string)(*[]df.Entry, error){
	file, err := os.Open(path)
	if err != nil{
		return nil, errors.New("cannot open file")
	}

	scanner := bufio.NewScanner(file)

	var entrySlice []df.Entry
	var entry df.Entry
	var splited []string
	var repoPath string
	var homePath string
	var line string

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")

	for scanner.Scan(){
		
		line = scanner.Text()

		if commentReg.MatchString(line) || emptyReg.MatchString(line){
			continue
		}

		splited = strings.Split(line, ":")
		repoPath = splited[0]
		homePath = splited[1]

		entry = df.NewEntry(repoPath, homePath)
		entrySlice = append(entrySlice, entry)
	}
	return &entrySlice,nil
}
