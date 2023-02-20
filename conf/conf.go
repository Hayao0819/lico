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

	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
)

// 設定ファイル全体
type List []Entry

func (item *Entry) String(replace bool) (string, error) {
	var (
		repo, home p.Path
	)
	var err error

	if replace {
		repo, err = replaceToTemplate(item.RepoPath.String())
		if err != nil {
			return "", err
		}

		home, err = replaceToTemplate(item.HomePath.String())
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
func (list *List) GetItemFromPath(path p.Path) *Entry {
	// Todo
	for _, item := range *list {
		//fmt.Printf("%v and %v, %v and %v\n", item.HomePath, path, item.RepoPath, path)
		if item.HomePath == path || item.RepoPath == path {
			return &item
		} else {
			homeIsSame, err := p.Is(item.HomePath, path)
			if err != nil {
				continue
			}
			repoIsSame, err := p.Is(item.RepoPath, path)
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

// パスがリポジトリファイルに含まれているかどうか
func (entries *List) HasRepoFile(path p.Path) (bool, error) {
	for _, entry := range *entries {
		result, err := p.Is(entry.RepoPath, path)
		if err != nil {
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
}

// パスがホームファイルに含まれているかどうか
func (entries *List) HasHomeFile(path p.Path) (bool, error) {
	for _, entry := range *entries {
		result, err := p.Is(entry.HomePath, path)
		if err != nil {
			continue
		}
		if result {
			return true, nil
		}
	}
	return false, nil
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
	var repoPath p.Path
	var homePath p.Path
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
		repoPath = p.Path(strings.TrimSpace(splited[0]))
		homePath = p.Path(strings.TrimSpace(splited[1]))

		item = NewEntryWithIndex(repoPath, homePath, lineNo)
		list = append(list, item)
	}
	return &list, nil
}

// テンプレートを解析してPathを生成します
func Format(path string) (p.Path, error) {
	var parsed p.Path

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

	parsed = p.Path(parsedBytes.String())

	return parsed, nil
}
