package conf

import (
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/vars"
)

// 設定ファイル全体
type List []Entry

// 指定されたパスを持つEntryを返します
func (list *List) GetItemFromPath(path p.Path) (*Entry, error) {
	path, err := path.Abs("")
	if err != nil{
		return nil, err
	}

	// Todo
	for _, item := range *list {
		//cmd.Printf("%v and %v, %v and %v\n", item.HomePath, path, item.RepoPath, path)
		if item.HomePath == path || item.RepoPath == path {
			return &item, nil
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
				return &item, nil
			}
		}
	}
	return nil, vars.ErrNoSuchEntry(path.String())
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
