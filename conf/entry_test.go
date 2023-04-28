package conf_test

import (
	"path"
	"testing"

	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/tester"
	"github.com/Hayao0819/lico/vars"
	p "github.com/Hayao0819/lico/paths"
)

func TestMain(m *testing.M){
	tester.CommonTestMain("../example")(m)
}

type testEntry struct{
	repo p.Path
	home p.Path
	index int
	expect conf.Entry
}

var args = []testEntry{
	{
		repo: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
		home: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
		index: 1,
		expect: conf.Entry{
			RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
			HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
		},
	},
	{
		repo: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
		home: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
		index: 2,
		expect: conf.Entry{
			RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
			HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
		},
	},
}

func getEntry() []conf.Entry{
	var entries []conf.Entry
	for _, arg := range args {
		entries = append(entries, arg.expect)
	}
	return entries
}

func TestNewEntry(t *testing.T){
	for _, arg := range args {
		entry := conf.NewEntry(arg.repo, arg.home)
		if entry != arg.expect {
			t.Errorf("conf.NewEntry(%s, %s) = %v", arg.repo, arg.home, entry)
		}
	}
}

func TestNewEntryWithIndex(t *testing.T){
	//tester.MakeSymLinkInExample()
	for _, arg := range args {
		entry := conf.NewEntryWithIndex(arg.repo, arg.home, arg.index)
		if entry != arg.expect {
			t.Errorf("conf.NewEntryWithIndex(%s, %s, %d) = %v", arg.repo, arg.home, arg.index, entry)
		}
	}
}

func TestExistsRepoPath(t *testing.T){
	type testcase struct{
		entry conf.Entry
		expect bool
	}

	tests := []testcase{}

	for _, e := range getEntry() {
		tests = append(tests, testcase{
			entry: e,
			expect: true,
		})
	}

	// add not exists path
	tests = append(tests, testcase{
		entry: conf.Entry{
			RepoPath: p.Path(path.Join(vars.RepoDir, "config", "not-exists.txt")),
			HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "not-exists.txt")),
		},
		expect: false,
	})

	for _, arg := range tests {
		if arg.entry.ExistsRepoPath() != arg.expect {
			t.Errorf("conf.ExistsRepoPath(%v) != %t", arg.entry, arg.expect)
		}
	}
}
