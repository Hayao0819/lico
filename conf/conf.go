package conf

import (
	"bufio"
	"text/template"
	//"errors"
	"os"
	"regexp"
	"strings"
	"bytes"

	"github.com/Hayao0819/lico/utils"
	df "github.com/Hayao0819/lico/dotfile"
)


func ReadConf(path string)(*[]df.Entry, error){
	file, err := os.Open(path)
	if err != nil{
		return nil,ErrCantOpenListFile
	}

	scanner := bufio.NewScanner(file)

	var entrySlice []df.Entry
	var entry df.Entry
	var splited []string
	var repoPath df.Path
	var homePath df.Path
	var line string

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")

	for scanner.Scan(){
		
		line = scanner.Text()

		if commentReg.MatchString(line) || emptyReg.MatchString(line){
			continue
		}

		splited = strings.Split(line, ":")
		repoPath = df.Path(splited[0])
		homePath = df.Path(splited[1])

		entry = df.NewEntry(repoPath, homePath)
		entrySlice = append(entrySlice, entry)
	}
	return &entrySlice,nil
}

func Format(path string)(df.Path, error){
	var parsed df.Path
	
	dirInfo, err := utils.GetOSEnv()
	if err != nil{
		return parsed, err
	}

	tpl, err := template.New("path").Parse(path)
	if err != nil{
		return parsed, err
	}
	var parsedBytes bytes.Buffer
	if err := tpl.Execute(&parsedBytes, dirInfo); err !=nil{
		return parsed, err
	}

	parsed = df.Path(parsedBytes.String())

	return parsed,nil
}

/*
func WriteEntries(entries *[]df.Entry, listFile string)(error){
	file, err := os.Create(listFile)
	if err != nil{
		return ErrCantOpenListFile
	}

}
*/
