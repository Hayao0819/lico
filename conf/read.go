package conf

import (
	"regexp"
	"strings"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	p "github.com/Hayao0819/lico/paths"
)



func ReadCreatedList() (*List, error) {
	path := vars.CreatedListFile

	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}

	var list List

	for lineNo, line := range lines {
		list = append(list, NewEntryWithIndex("", p.Path(strings.TrimSpace(line)), lineNo+1))
	}

	return &list, nil
}

// 設定ファイルを読み込みます
func ReadConf() (*List, error) {
	path := vars.BaseListFile

	// parse config
	lines, err := FormatTemplate(path)
	if err != nil {
		return nil, err
	}

	var list List
	var item Entry
	var splited []string
	var repoPath p.Path
	var homePath p.Path

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")

	for lineNo, line := range lines {
		// コメントと空行を除外
		if commentReg.MatchString(line) || emptyReg.MatchString(line) {
			continue
		}

		// :で分割
		splited = strings.Split(line, ":")
		repoPath = p.Path(strings.TrimSpace(splited[0]))
		homePath = p.Path(strings.TrimSpace(splited[1]))

		//fmt.Println(repoPath+"=="+homePath)

		item = NewEntryWithIndex(repoPath, homePath, lineNo+1)
		list = append(list, item)
	}
	return &list, nil
}
