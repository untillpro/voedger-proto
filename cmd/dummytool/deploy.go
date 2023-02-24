package main

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
	"github.com/untillpro/goutils/logger"
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
				domain, ip, ok := validateNodeAddr(arg)
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
			dryRun, _ := cmd.Root().PersistentFlags().GetBool("dry-run")
			logger.Verbose("dry-run:", dryRun)

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

func validateNodeAddr(nodeAddr string) (domain string, ip net.IP, valid bool) {

	if strings.Contains(nodeAddr, ":") {
		parts := strings.Split(nodeAddr, ":")
		if len(parts) != 2 {
			return "", nil, false
		}
		domain = parts[0]
		ip = net.ParseIP(parts[1])
		if ip == nil {
			return "", nil, false
		}
		return
	}

	ip = net.ParseIP(nodeAddr)
	valid = ip != nil
	return
}
