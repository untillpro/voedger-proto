package main

import (
	_ "embed"

	// "fmt"
	"os"

	"github.com/untillpro/goutils/cobrau"
)

//go:embed version
var version string

var verbose bool

func main() {
	if err := rootCmd(os.Args, version); err != nil {
		os.Exit(1)
	}
}

func rootCmd(args []string, ver string) error {
	return cobrau.PrepareAndExecuteRootCmd(
		"ctool",
		"Cluster management utility",
		args, &version, &verbose,
		newDeployCmd(), newUpgradeCmd(),
	)
}
