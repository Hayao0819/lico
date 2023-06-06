package conf_test

import (
	"github.com/Hayao0819/lico/conf"
	"testing"
	//"github.com/Hayao0819/lico/vars"
)

func TestReadIgnoreList(t *testing.T) {
	ignore := conf.ReadIgnoreList()
	if ignore == nil {
		t.Errorf("conf.ReadIgnoreList() = nil")
	}
}
