/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/untillpro/goutils/logger"
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
			for i := 0; i < 10 && cmd.Context().Err() == nil; i++ {
				fmt.Println("Hello, world, I'm doing upgrade:", i, "!")
				time.Sleep(time.Second)
			}
			if cmd.Context().Err() != nil {
				logger.Verbose("upgrade command is interrupted")
			}

		},
	}
}
