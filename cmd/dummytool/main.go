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

	rootCmd := cobrau.PrepareRootCmd(
		"ctool",
		"Cluster management utility",
		args,
		newDeployCmd(), newUpgradeCmd(),
	)

	// Can be got as cmd.Root().PersistentFlags().GetBool("dry-run")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Simulate the execution of the command without actually modifying any files or data")

	return rootCmd.Execute()
}
