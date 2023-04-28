package vars

import "os"

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

func init(){
	if version == "Unknown" && commit == "Unknown" && date == "Unknown" {
		println("Please compile with goreleaser.")
		os.Exit(1)
	}
}
