/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package cobrau

import (
	"github.com/spf13/cobra"
	"github.com/untillpro/goutils/logger"
)

func PrepareAndExecuteRootCmd(use string, short string, args []string, version *string, verbose *bool, cmds ...*cobra.Command) error {

	var rootCmd = &cobra.Command{
		Use:   use,
		Short: "Cluster management utility written in golang",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if *verbose {
				logger.SetLogLevel(logger.LogLevelVerbose)
				logger.Verbose("Running in verbose mode...")
			}
		},
	}

	rootCmd.SetArgs(args[1:])
	rootCmd.AddCommand(cmds...)
	rootCmd.PersistentFlags().BoolVarP(verbose, "verbose", "v", false, "Print verbose output")
	rootCmd.SilenceUsage = true
	return rootCmd.Execute()
}
