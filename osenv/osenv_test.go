package osenv_test

import (
	"strings"
	"testing"

	"github.com/Hayao0819/lico/osenv"
	"github.com/Hayao0819/lico/tester"
)

func TestMain(m *testing.M) {
	tester.CommonTestMain("../example")(m)
}

func TestGet(t *testing.T) {
	e, err := osenv.Get()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range e {
		t.Logf("%s: %s", i, v)
		if strings.TrimSpace(i) == "" {
			t.Errorf("empty key: %s", i)
			break
		}
		if strings.TrimSpace(v) == "" {
			t.Errorf("empty value: %s = %s", i, v)
			break
		}
	}
}
