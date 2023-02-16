package conf

import (
	"fmt"
	"strings"

	df "github.com/Hayao0819/lico/dotfile"
	"github.com/Hayao0819/lico/utils"
)


func ReplaceToTemplate(path string)(df.Path, error){
	var parsed df.Path
	dirInfo, err := utils.GetOSEnv()
	if err != nil{
		return parsed, err
	}

	for _, key := range dirInfo.GetKeys(){
		path = strings.ReplaceAll(path, dirInfo[key], fmt.Sprintf("{{ .%v }}", key))
	}

	parsed = df.NewPath(path)
	return parsed, nil
}

