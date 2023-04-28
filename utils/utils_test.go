package utils_test

import (
	"testing"

	//	"github.com/Haya0819/lico/tester"
	"github.com/Hayao0819/lico/utils"
)


func TestSortWithLen(t *testing.T){
	args := []struct{
		input []string
		expect []string
	}{
		{
			input: []string{"cat", "apple", "dog", "banana", "elephant", "ant", "bird"},
			expect: []string{"elephant", "banana", "apple", "bird", "cat", "dog", "ant"},
		},
		{
			input: []string{"sayaka", "kyoko", "mami", "nagisa", "homura", "madoka"},
			expect: []string{"sayaka", "nagisa", "homura", "madoka", "kyoko", "mami"},

		},
	}
	for _, arg := range args {
		sorted := utils.SortWithLen(arg.input)
		if len(sorted) != len(arg.expect){
			t.Errorf("utils.SortWithLen(%v) = %v", sorted, arg.expect)
			break
		}
		for i, v := range sorted{
			if v != arg.expect[i]{
				t.Errorf("utils.SortWithLen(%v) = %v", sorted, arg.expect)
				break
			}
		}
	}
}

func TestGetHomeDir(t *testing.T){
	home := utils.GetHomeDir()
	if home == ""{
		t.Errorf("utils.GetHomeDir() = %s", home)
	}
}
