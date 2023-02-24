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
	return &cobra.Command{
		Use:   "deploy [[<domain>:]<ipaddr>...]",
		Short: "Deploys a cluster using specified nodes",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			var errArgs error
			for idx, arg := range args {
				domain, ip, ok := validateNodeAddr(arg)
				if !ok {
					errArgs = errors.Join(errArgs, fmt.Errorf("%w: actual argument #%v: %v", ErrDeployInvalidArg, idx, arg))
				}
				logger.Verbose("domain, ip", domain, ip)
			}
			if errArgs != nil {
				return errArgs
			}
			// TODO: Implement the deploy functionality using appCompose and dbCompose
			return nil

		},
	}
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
		addrs, err := net.LookupHost(domain)
		if err != nil {
			return "", nil, false
		}
		for _, addr := range addrs {
			if addr == ip.String() {
				return domain, ip, true
			}
		}
		return "", nil, false
	}

	ip = net.ParseIP(nodeAddr)
	if ip != nil {
		return "", ip, true
	}

	return "", nil, false
}
