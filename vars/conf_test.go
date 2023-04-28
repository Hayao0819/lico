package vars_test

import (
	"log"
	"os"
	"testing"

	"github.com/Hayao0819/lico/vars"
)

func TestEnableTestMode(t *testing.T) {
	vars.EnableTestMode("../example")
	if _, err := os.Stat(vars.RepoDir); os.IsNotExist(err) {
		t.Errorf("vars.RepoDir is not exist")
	}else{
		log.Println("vars.RepoDir is exist:", vars.RepoDir)
	}
}
