package pkglist_test

import (
	"log"
	"testing"

	"github.com/Hayao0819/lico/pkglist"
	"github.com/Hayao0819/lico/tester"
)

func TestMain(m *testing.M) {
	tester.CommonTestMain("../example")(m)
}

func TestReadList(t *testing.T) {
	list, err := pkglist.ReadList()
	if err != nil {
		t.Errorf("pkglist.ReadList() error: %v", err)
	} else {
		log.Printf("List: %v\n", list)
	}
}

func TestOSList(t *testing.T) {
	list, err := pkglist.ReadList()
	if err != nil {
		t.Errorf("pkglist.ReadList() error: %v", err)
	}

	oslist := list.OSList()
	if len(oslist) == 0 {
		t.Errorf("pkglist.OSList() error: %v", err)
	} else {
		log.Printf("OSList: %v\n", oslist)
	}
}

func TestGetCurrent(t *testing.T) {
	list, err := pkglist.ReadList()
	if err != nil {
		t.Fatal(err)
	}
	current, err := list.GetCurrent()
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Current: %v\n", current)
}
