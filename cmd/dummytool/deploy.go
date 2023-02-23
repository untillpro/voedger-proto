package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDeployCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [IP or Domain:IP] [IP or Domain:IP] [IP or Domain:IP] [IP or Domain:IP] [IP or Domain:IP]",
		Short: "Deploy command deploys a cluster using specified nodes",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if !isValidAddress(arg) {
					return fmt.Errorf("invalid argument - %s\nPlease provide arguments in the format - IP or Domain:IP", arg)
				}
			}
			// TODO: Implement the deploy functionality using appCompose and dbCompose
			return nil
		},
	}
}
