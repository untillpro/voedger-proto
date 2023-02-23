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
	if ip == nil {
		return false
	}
	return true
}

//go:embed version
var version string

func main() {
	if err := execute(os.Args[1:], version); err != nil {
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

	rootCmd.SetArgs(args)
	rootCmd.AddCommand(newDeployCmd(), newUpgradeCmd())
	return rootCmd.Execute()
}
