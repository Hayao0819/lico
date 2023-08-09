package tester_test

import (
	"os"
	"testing"

	"github.com/Hayao0819/lico/tester"
	"github.com/Hayao0819/lico/vars"
)

func TestEnable(t *testing.T) {
	err := tester.Enable("../example")
	if err != nil {
		t.Fatal(err)
	}

	if _, err = os.Stat(vars.RepoDir); err != nil {
		t.Fatal(err)
	}
}

func TestCommonTestMain(t *testing.T) {
	f := tester.CommonTestMain("../example")
	if f == nil {
		t.Fatal("nil function")
	}
}
