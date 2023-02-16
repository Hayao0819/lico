package conf

import (
	"bufio"
	"fmt"
	"text/template"

	//"errors"
	"bytes"
	"os"
	"regexp"
	"strings"

	df "github.com/Hayao0819/lico/dotfile"
	"github.com/Hayao0819/lico/utils"

	
)

//import errList "github.com/Hayao0819/lico/errlist"

// 設定ファイルの1行であるdotfile.Entryに行番号を追加したもの
type ListItem struct{
	Entry df.Entry
	Index int
}

// 
func NewListItem(entry df.Entry)(ListItem){
	return ListItem{
		Entry:  entry,
		Index: 0,
	}
}

func NewListItemWithIndex(entry df.Entry, index int)(ListItem){
	return ListItem{
		Entry:  entry,
		Index: index,
	}
}


// 設定ファイル全体
type List []ListItem

// 設定ファイル全体からEntryを全て取り出します
func (list *List)GetEntries()(*[]df.Entry){
	var rtn []df.Entry 
	for _,listitem := range *list{
		rtn = append(rtn, listitem.Entry)
	}
	return &rtn
}


func (item *ListItem) String ()(string, error){
	repo, err := ReplaceToTemplate(item.Entry.RepoPath.String())
	if err !=nil{
		return "", err
	}

	home, err := ReplaceToTemplate(item.Entry.HomePath.String())
	if err !=nil{
		return "", err
	}

	return fmt.Sprintf("%v:%v\n", repo, home), nil
}

// 指定されたパスを持つListItemを返します
func (list *List)GetItemFromPath(path df.Path)(*ListItem){
	// Todo
	for _, item := range *list{
		fmt.Printf("%v and %v, %v and %v\n", item.Entry.HomePath, path, item.Entry.RepoPath, path)
		if item.Entry.HomePath == path || item.Entry.RepoPath == path{
			return &item
		}else{
			homeIsSame, err := df.PathIs(item.Entry.HomePath, path)
			if err!=nil{
				continue
			}
			repoIsSame , err := df.PathIs(item.Entry.RepoPath, path)
			if err!=nil{
				continue
			}
			if homeIsSame || repoIsSame{
				return &item
			}
		}
	}
	return nil
}

// 設定ファイルを読み込みます
func ReadConf(path string)(*List, error){
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	

	var list List
	var item ListItem
	var splited []string
	var repoPath df.Path
	var homePath df.Path
	var line string

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")


	lineNo := 0
	for scanner.Scan(){
		lineNo++
		line = scanner.Text()

		if commentReg.MatchString(line) || emptyReg.MatchString(line){
			continue
		}

		splited = strings.Split(line, ":")
		repoPath = df.Path(strings.TrimSpace(splited[0]))
		homePath = df.Path(strings.TrimSpace(splited[1]))

		item = NewListItemWithIndex(df.NewEntry(repoPath, homePath), lineNo) 
		list = append(list, item)
	}
	return &list,nil
}

// テンプレートを解析してPathを生成します
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

