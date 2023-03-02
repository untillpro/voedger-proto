/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package cobrau

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/untillpro/goutils/logger"
)

func ExecCommandAndCatchInterrupt(cmd *cobra.Command) error {

	ctxDone, cancel := context.WithCancel(context.Background())

	cmdExec := func(ctx context.Context) (err error) {
		err = cmd.ExecuteContext(ctx)
		cancel()
		return
	}

	return goAndCatchInterrupt(cmdExec, ctxDone)

}

func goAndCatchInterrupt(f func(ctx context.Context) error, ctxDone context.Context) (err error) {

	var signals = make(chan os.Signal, 1)

	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(signals, os.Interrupt)

	go func() {
		err = f(ctx)
	}()

	for ctxDone.Err() == nil {
		select {
		case sig := <-signals:
			if ctx.Err() == nil {
				logger.Verbose("signal received:", sig)
				cancel()
			}
		case <-ctxDone.Done():
			logger.Verbose("ctxDone closed")
		}
	}
	cancel()
	return
}
