package conf_test

import (
	"path"
	"testing"

	"github.com/Hayao0819/lico/conf"
	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/tester"
	"github.com/Hayao0819/lico/vars"
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

func TestNewEntry(t *testing.T){
	var args = []testEntry{
		{
			repo: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
			home: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
			expect: conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
			},
		},
		{
			repo: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
			home: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
			expect: conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
			},
		},
	}


	for _, arg := range args {
		entry := conf.NewEntry(arg.repo, arg.home)
		if entry != arg.expect {
			t.Errorf("conf.NewEntry(%s, %s) = %v", arg.repo, arg.home, entry)
		}
	}
}

func TestNewEntryWithIndex(t *testing.T){
	var args = []testEntry{
		{
			repo: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
			home: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
			index: 1,
			expect: conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
				Index: 1,
			},
		},
		{
			repo: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
			home: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
			index: 2,
			expect: conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
				Index: 2,
			},
		},
	}
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

	tests := []testcase{
		{
			entry:  conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example2.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex2.txt")),
			},
			expect: true,
		},
		{
			entry: conf.Entry{
				RepoPath: p.Path(path.Join(vars.RepoDir, "config", "example1.txt")),
				HomePath: p.Path(path.Join(vars.RepoDir, "your-home", "ex1.txt")),
			},
			expect: true,
		},
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
