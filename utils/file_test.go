package utils_test

import (
	//"os"
	"testing"

	"github.com/Hayao0819/lico/tester"
	"github.com/Hayao0819/lico/utils"
)

/*
func makeTestCase() error{
	tempdir, err := os.MkdirTemp("", "licotest")
	testcase := []struct{
		path string
		content string
	}

}
*/

func TestReadLines(t *testing.T){
	testcase := []struct{
		path string
		want []string
	}{
		{
			path: tester.RepoRoot +  "/testdata/test.txt",
			want: []string{"test1", "test2", "test3"},
		},
	}

	for _, tc := range testcase{
		got, err := utils.ReadLines(tc.path)
		if err != nil{
			t.Errorf("ReadLines() error: %v", err)
		}
		for i := range got{
			if got[i] != tc.want[i]{
				t.Errorf("ReadLines() got: %v, want: %v", got[i], tc.want[i])
			}
		}
	}
}
