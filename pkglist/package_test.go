package pkglist_test

import (
	"testing"
	"github.com/Hayao0819/lico/pkglist"
)

func TestNewPkg(t *testing.T){
	p := pkglist.NewPkg("test")
	if p != "test" {
		t.Errorf("NewPkg(\"test\") = %v, want \"test\"", p)
	}
}
