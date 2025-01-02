/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import  (
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/ocular-d/wingnut/cmd"
)

// These values are automatically populated by Goreleaser
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)


func main() {
	// if installing from source, i.e. not a binary built by `Goreleaser`, this version will not be updated. In that case, use the automatic detection in `versioninfo`
	if version == "dev" {
		version = versioninfo.Version
		commit = versioninfo.Revision
		date = versioninfo.LastCommit.Format(time.RFC3339)
	} else {
		// Goreleaser doesn't prefix with a `v`, which we expect
		version = "v" + version
	}
	
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
