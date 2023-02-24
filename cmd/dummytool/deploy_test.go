package main

import (
	"testing"

	"github.com/untillpro/goutils/testingu"
)

func TestBasicUsage_Deploy(t *testing.T) {

	testCases := []testingu.RootTestCase{
		{
			Name:               "Wrong number of arguments",
			Args:               []string{"dummyutil", "deploy", "1.0.0.1"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        nil,
			ExpectedErrPattern: "5 arg",
		},
		{
			Name:               "Invalid argument",
			Args:               []string{"dummyutil", "deploy", "1", "2", "3", "4", "5"},
			Version:            "0.0.0-dummy",
			ExpectedErr:        ErrDeployInvalidArg,
			ExpectedErrPattern: "invalid argument",
		},
	}

	testingu.RunRootTestCases(t, rootCmd, testCases)

}
