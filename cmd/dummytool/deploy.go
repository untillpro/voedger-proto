package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDeployCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy {[<domain>:]<ipaddr>}",
		Short: "Deploy command deploys a cluster using specified nodes",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			for idx, arg := range args {
				if !isValidAddress(arg) {
					return fmt.Errorf("%w: argument #%v: %v", ErrDeployInvalidArg, idx, arg)
				}
			}
			// TODO: Implement the deploy functionality using appCompose and dbCompose
			return nil
		},
	}
}
