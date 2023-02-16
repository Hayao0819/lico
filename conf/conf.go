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

	"github.com/Hayao0819/lico/utils"
)

// 設定ファイル全体
type List []Entry

func (item *Entry) String(replace bool) (string, error) {
	var (
		repo, home Path
	)
	var err error

	if replace {
		repo, err = ReplaceToTemplate(item.RepoPath.String())
		if err != nil {
			return "", err
		}

		home, err = ReplaceToTemplate(item.HomePath.String())
		if err != nil {
			return "", err
		}
	} else {
		repo = item.RepoPath
		home = item.HomePath
	}
	return fmt.Sprintf("%v:%v\n", repo, home), nil
}

// 指定されたパスを持つEntryを返します
func (list *List) GetItemFromPath(path Path) *Entry {
	// Todo
	for _, item := range *list {
		fmt.Printf("%v and %v, %v and %v\n", item.HomePath, path, item.RepoPath, path)
		if item.HomePath == path || item.RepoPath == path {
			return &item
		} else {
			homeIsSame, err := PathIs(item.HomePath, path)
			if err != nil {
				continue
			}
			repoIsSame, err := PathIs(item.RepoPath, path)
			if err != nil {
				continue
			}
			if homeIsSame || repoIsSame {
				return &item
			}
		}
	}
	return nil
}

// 設定ファイルを読み込みます
func ReadConf(path string) (*List, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list List
	var item Entry
	var splited []string
	var repoPath Path
	var homePath Path
	var line string

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")

	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line = scanner.Text()

		if commentReg.MatchString(line) || emptyReg.MatchString(line) {
			continue
		}

		splited = strings.Split(line, ":")
		repoPath = Path(strings.TrimSpace(splited[0]))
		homePath = Path(strings.TrimSpace(splited[1]))

		item = NewEntryWithIndex(repoPath, homePath, lineNo)
		list = append(list, item)
	}
	return &list, nil
}

// テンプレートを解析してPathを生成します
func Format(path string) (Path, error) {
	var parsed Path

	dirInfo, err := utils.GetOSEnv()
	if err != nil {
		return parsed, err
	}

	tpl, err := template.New("path").Parse(path)
	if err != nil {
		return parsed, err
	}
	var parsedBytes bytes.Buffer
	if err := tpl.Execute(&parsedBytes, dirInfo); err != nil {
		return parsed, err
	}

	parsed = Path(parsedBytes.String())

	return parsed, nil
}
