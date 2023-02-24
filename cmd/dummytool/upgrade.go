/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newUpgradeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: fmt.Sprintf("Update the cluster components to version %s if necessary", version),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				fmt.Println("Error: upgrade command does not take any arguments")
				os.Exit(1)
			}
			// TODO: Implement the upgrade functionality
		},
	}
}
