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
			logger.Verbose("Deploying SE")
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
			// TODO: Implement the deploy SE functionality
			return nil
		},
	}
	deployCECmd := &cobra.Command{
		Use:   "CE [<ipaddr>]",
		Short: "Deploy CE on the specified node",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Verbose("Deploying CE")
			// TODO: Implement the deploy CE functionality
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
