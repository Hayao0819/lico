package conf

import (
	"github.com/Hayao0819/lico/vars"
	gi "github.com/sabhiram/go-gitignore"
)

type IgnoreList gi.GitIgnore

// lico.ignoreを読み込んでIgnoreListを生成する
func ReadIgnoreList() (*IgnoreList) {
	lines := []string{}
	for _, i := range vars.GetIgnore(){
		_lines, err := FormatTemplate(i)
		if err != nil {
			continue
		}
		lines = append(lines, _lines...)
	}

	gitignore := gi.CompileIgnoreLines(lines...)

	return (*IgnoreList)(gitignore)
}

// パスがIgnoreListに含まれているかどうか
func (i *IgnoreList) MatchString(s string) (bool, string) {
	g := gi.GitIgnore(*i)
	b, h := g.MatchesPathHow(s)
	if h == nil {
		return b, ""
	}
	return b, h.Line
}

// entryがIgnoreListに含まれているかどうか
func (i *IgnoreList) MatchEntry(e Entry) (bool, string) {
	homeStatus, homeStr := i.MatchString(e.HomePath.String())
	repoStatus, repoStr := i.MatchString(e.RepoPath.String())
	if homeStatus {
		return true, homeStr
	}
	if repoStatus {
		return true, repoStr
	}
	return false, ""
}
