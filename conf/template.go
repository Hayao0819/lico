package conf

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Hayao0819/lico/osenv"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"os"
)

func replaceToTemplate(path string) (p.Path, error) {
	var parsed p.Path
	dirInfo, err := osenv.Get()
	if err != nil {
		return parsed, err
	}

	for _, key := range dirInfo.GetKeys() {
		path = strings.ReplaceAll(path, dirInfo[key], fmt.Sprintf("{{ .%v }}", key))
	}

	parsed = p.New(path)
	return parsed, nil
}


// テンプレートを解析してPathを生成します
func FormatTemplate(path string) ([]string, error) {
	parsed := []string{}

	dirInfo, err := osenv.Get()
	if err != nil {
		return parsed, err
	}

	funcMap := template.FuncMap{
		"environ": func(n string) string {
			return os.Getenv(n)
		},
		"isempty": func(s string) bool {
			return utils.IsEmpty(s)
		},
		"isset": func(key string) bool {
			return !utils.IsEmpty(os.Getenv(key))
		},
		"is_installed": func (c string)bool{
			_, s := os.LookupEnv(c)
			return s
		},
		"isunix": func()bool{
			return dirInfo.Get("OS") == "linux" || dirInfo.Get("OS") == "darwin"
		},
	}
	

	tpl, err := template.New(filepath.Base(path)).Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return parsed, err
	}

	var parsedBytes bytes.Buffer
	if err := tpl.Execute(&parsedBytes, dirInfo); err != nil {
		return parsed, err
	}

	parsed = strings.Split(parsedBytes.String(), "\n")

	return parsed, nil
}
