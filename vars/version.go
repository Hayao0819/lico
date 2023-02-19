package vars

type verInfo struct {
	Name   string
	Commit string
	Date   string
}

var version string = "Unknown"
var commit string = "Unknown"
var date string = "Unknown"

var Version = verInfo{
	Name:   version,
	Commit: commit,
	Date:   date,
}
