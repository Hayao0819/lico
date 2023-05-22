package vars_test

import (
	"testing"

	"github.com/Hayao0819/lico/vars"
)
func TestGetRepoDir(t *testing.T){
	if vars.GetRepoDir() != vars.RepoDir{
		t.Fatal("GetRepoDir() != vars.RepoDir")
	}
}

func TestCreated(t *testing.T){
	if vars.GetCreated() != vars.Created{
		t.Fatal("GetCreated() != vars.Created")
	}
}
