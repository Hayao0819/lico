package pkglist_test

import (
	"github.com/Hayao0819/lico/pkglist"
	"testing"
)

func TestNewPkg(t *testing.T) {
	p := pkglist.NewPkg("test")
	if p != "test" {
		t.Errorf("NewPkg(\"test\") = %v, want \"test\"", p)
	}
}
