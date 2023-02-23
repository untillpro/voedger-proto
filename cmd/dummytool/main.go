package main

import (
	_ "embed"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func isValidAddress(address string) bool {
	// Address should be in the format of IP or Domain:IP
	parts := strings.Split(address, ":")
	if len(parts) != 2 {
		return false
	}
	ip := net.ParseIP(parts[1])
	return !(ip == nil)
}

//go:embed version
var version string

func main() {
	if err := execute(os.Args, version); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func execute(args []string, ver string) error {

	version = ver

	var rootCmd = &cobra.Command{
		Use:   "ctool",
		Short: "Cluster management utility written in golang",
	}

	rootCmd.SetArgs(args[1:])
	rootCmd.AddCommand(newDeployCmd(), newUpgradeCmd())
	return rootCmd.Execute()
}
