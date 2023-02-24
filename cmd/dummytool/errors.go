package main

import "errors"

// nolint
var (
	ErrDeployInvalidArg = errors.New("invalid argument format, expected <ipaddr> or <domain>:<ipaddr>")
)
