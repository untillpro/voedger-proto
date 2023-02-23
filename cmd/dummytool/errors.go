package main

import "errors"

// nolint
var (
	ErrDeployInvalidArg = errors.New("invalid argument, use IP or Domain:IP")
)
