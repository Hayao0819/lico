package conf

import (
	"fmt"
	"strings"

	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
)

func ReplaceToTemplate(path string) (p.Path, error) {
	var parsed p.Path
	dirInfo, err := utils.GetOSEnv()
	if err != nil {
		return parsed, err
	}

	for _, key := range dirInfo.GetKeys() {
		path = strings.ReplaceAll(path, dirInfo[key], fmt.Sprintf("{{ .%v }}", key))
	}

	parsed = p.New(path)
	return parsed, nil
}
