package tester

import (
	"os"
	"path"
	"testing"

	"github.com/Hayao0819/lico/vars"
)

// テストモードを有効にする
func Enable(relToExample string)error{
	current_dir , err := os.Getwd()
	if err !=nil{
		return err
	}
	vars.RepoDir = path.Join(current_dir , relToExample)
	return nil
}

func CommonTestMain(relToExample string)(func(*testing.M)){
	return func(m *testing.M){
		if err := Enable(relToExample); err != nil{
			os.Exit(-1)
		}
		os.Exit(m.Run())
	}
}
