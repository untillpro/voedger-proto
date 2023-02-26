/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/untillpro/goutils/logger"
	"voedger.io/voedger/cmd/dummytool/internal/mynet"
)

func newDeployCmd() *cobra.Command {
	deploySECmd := &cobra.Command{
		Use:   "SE [[<domain>:]<ipaddr>...]",
		Short: "Deploy an SE cluster using the specified nodes",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Verbose("deploying SE")
			var errArgs error
			for idx, arg := range args {
				domain, ip, ok := mynet.ValidateNodeAddr(arg)
				if !ok {
					errArgs = errors.Join(errArgs, fmt.Errorf("%w: actual argument #%v: %v", ErrDeployInvalidArg, idx+1, arg))
				}
				if logger.IsVerbose() { // you might skip this check if perfomance is not your goal
					logger.Verbose("domain, ip", domain, ip)
				}
			}
			if errArgs != nil {
				return errArgs
			}
			dryRun, _ := cmd.Root().PersistentFlags().GetBool("dry-run")
			if dryRun {
				logger.Verbose("I'm in dry-run mode")
				return nil
			}
			logger.Verbose("normal mode")
			return nil
		},
	}
	deployCECmd := &cobra.Command{
		Use:   "CE [<ipaddr>]",
		Short: "Deploy CE on the specified node",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Verbose("Deploying CE")

			dryRun, _ := cmd.Root().PersistentFlags().GetBool("dry-run")
			if dryRun {
				logger.Verbose("I'm in dry-run mode")
				return nil
			}
			logger.Verbose("normal mode")
			return nil
		},
	}

	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy CE or SE cluster",
	}
	deployCmd.AddCommand(deployCECmd, deploySECmd)

	return deployCmd
}
