package conf

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"os"

	"github.com/Hayao0819/go-distro/goos"
	"github.com/Hayao0819/lico/osenv"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
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

func GetTemplateFuncMap() *template.FuncMap {
	funcMap := template.FuncMap{
		"environ": func(n string) string {
			return os.Getenv(n)
		},
		"is_empty": func(s string) bool {
			return utils.IsEmpty(s)
		},
		"is_set": func(v string) bool {
			//return !utils.IsEmpty(os.Getenv(key))
			_, s := os.LookupEnv(v)
			return s
		},
		"is_installed": func(c string) bool {
			_, err := exec.LookPath(c)
			if err == nil {
				return true
			} else {
				return false
			}
		},
		"is_unix": func() bool {
			return goos.IsUnix()
		},
		"readdir": func(p string) []string {
			direntry, err := os.ReadDir(p)
			if err != nil {
				return []string{}
			}
			files := []string{}
			for _, entry := range direntry {
				files = append(files, entry.Name())
			}
			return files
		},
		"readdir_files": func(p string) []string {
			direntry, err := os.ReadDir(p)
			if err != nil {
				return []string{}
			}
			files := []string{}
			for _, entry := range direntry {
				if entry.IsDir() {
					continue
				}
				files = append(files, entry.Name())
			}
			return files
		},
		"joinpath": func(p ...string) string {
			return filepath.Join(p...)
		},
		"is_global": func() bool {
			return vars.GlobalMode
		},
		"is_exist": func(path string) bool {
			return utils.Exists(path)
		},
		"is_systemd_running": func() bool {
			cmd := exec.Command("systemctl", "status")
			/*
				if cmd.Run() == nil {
					return true
				}
				return false
			*/
			return cmd.Run() == nil
		},
		// Todo
		//"is_systemd_service_enabled": func (service string)bool{
		//},
		//"is_systemd_service_active": func(service string)bool{
		//},
	}

	return &funcMap
}

// テンプレートを解析してPathを生成します
func FormatTemplate(path string) ([]string, error) {
	parsed := []string{}

	dirInfo, err := osenv.Get()
	if err != nil {
		return parsed, err
	}

	funcMap := *GetTemplateFuncMap()
	funcMap["isempty"] = funcMap["is_empty"]
	funcMap["isset"] = funcMap["is_set"]
	funcMap["isunix"] = funcMap["is_unix"]
	funcMap["isglobal"] = funcMap["is_global"]

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
