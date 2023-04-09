package conf

import (
	"regexp"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)


type IgnoreList []*regexp.Regexp

// lico.ignoreを読み込んでIgnoreListを生成する
func ReadIgnoreList()(*IgnoreList, error){
	lines, err := FormatTemplate(vars.IgnoreListFile)
	if err !=nil{
		return nil, err
	}

	exps := IgnoreList{}

	for _, line := range lines{
		if utils.IsEmpty(line){
			continue
		}
		e, err := regexp.Compile(line)
		if err !=nil{
			return nil, err
		}
		exps = append(exps, e)
	}
	return &exps, nil
}

// パスがIgnoreListに含まれているかどうか
func (i *IgnoreList)MatchString(s string)(bool, string){
	for _, r := range *i {
		if r.MatchString(s) {
			return true, r.String()
		}
	}
	return false, ""
}

// entryがIgnoreListに含まれているかどうか
func (i *IgnoreList)MatchEntry(e Entry)(bool, string){
	homeStatus, homeStr := i.MatchString(e.HomePath.String())
	repoStatus, repoStr := i.MatchString(e.RepoPath.String())
	if homeStatus{
		return true, homeStr
	}
	if repoStatus{
		return true, repoStr
	}
	return false , ""
}
