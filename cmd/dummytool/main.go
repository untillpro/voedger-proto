package main

import (
	_ "embed"

	// "fmt"
	"os"

	"github.com/untillpro/goutils/cobrau"
)

//go:embed version
var version string

func main() {
	if err := execRootCmd(os.Args, version); err != nil {
		os.Exit(1)
	}
}

func execRootCmd(args []string, ver string) error {
	version = ver
	return cobrau.PrepareRootCmd(
		"ctool",
		"Cluster management utility",
		args,
		newDeployCmd(), newUpgradeCmd(),
	).Execute()
}
