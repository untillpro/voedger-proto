/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import "errors"

// nolint
var (
	ErrDeployInvalidArg = errors.New("invalid argument format, expected <ipaddr> or <domain>:<ipaddr>")
)
