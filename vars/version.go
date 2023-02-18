package vars

type verInfo struct{
	Name string
	Commit string
	Date string
}


var version, commit, date string

var Version = verInfo{
	Name: version,
	Commit: commit,
	Date: date,
}
