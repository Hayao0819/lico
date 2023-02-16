package conf

import (
	"fmt"
	"strings"

	"github.com/Hayao0819/lico/utils"
)

func ReplaceToTemplate(path string) (Path, error) {
	var parsed Path
	dirInfo, err := utils.GetOSEnv()
	if err != nil {
		return parsed, err
	}

	for _, key := range dirInfo.GetKeys() {
		path = strings.ReplaceAll(path, dirInfo[key], fmt.Sprintf("{{ .%v }}", key))
	}

	parsed = NewPath(path)
	return parsed, nil
}
