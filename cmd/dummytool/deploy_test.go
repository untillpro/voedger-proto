/*
* Copyright (c) 2023-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import (
	"testing"

	"github.com/untillpro/goutils/testingu"
)

func TestBasicUsage_Deploy(t *testing.T) {

	testCases := []testingu.RootTestCase{
		{
			Name:               "Wrong number of arguments",
			Args:               []string{"dummyutil", "deploy", "SE", "1.0.0.1"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        nil,
			ExpectedErrPattern: "5 arg(s)",
		},
		{
			Name:               "Invalid argument",
			Args:               []string{"dummyutil", "deploy", "SE", "1", "2", "3", "4", "5"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        ErrDeployInvalidArg,
			ExpectedErrPattern: "invalid argument",
		},
		{
			Name:               "Good args, dry-run",
			Args:               []string{"dummyutil", "deploy", "SE", "--dry-run", "1.1.1.1", "1.1.1.2", "1.1.1.3", "1.1.1.4", "1.1.1.5"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        nil,
			ExpectedErrPattern: "",
		},
		{
			Name:               "Good args, dry-run, IPv6",
			Args:               []string{"dummyutil", "deploy", "SE", "--dry-run", "fe80::215:5dff:fe6b:3d02", "192.1.1.2", "1.1.1.3", "1.1.1.4", "1.1.1.5"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        nil,
			ExpectedErrPattern: "",
		},
	}

	testingu.RunRootTestCases(t, execRootCmd, testCases)

}
