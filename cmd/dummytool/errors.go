package main

import "errors"

// nolint
var (
	ErrDeployInvalidArg = errors.New("invalid argument, use <ipaddr> or <domain>:<ipaddr>")
)
