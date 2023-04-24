package vars

var RepoDir string = "~/.lico/repo"
var BaseListFile string = "~/.lico/repo/lico.list"
var IgnoreListFile string = "~/.lico/repo/lico.ignore"
var CreatedListFile string = "~/.lico/created.list"
var PkgListFile string = "~/.lico/repo/lico-pkgs-2.json"

var HomeDir string
//var RepoPathBase, HomePathBase *string

var (
	RepoPathBase *string = &RepoDir
	HomePathBase *string = &HomeDir
)


