package main

import (
	"fmt"
	"testing"

	"github.com/untillpro/goutils/tests"
)

func TestBasicUsage_Deploy(t *testing.T) {

	testCases := []struct {
		name               string
		args               []string
		version            string
		expectedErr        error
		expectedErrPattern string
	}{
		{
			name:               "Wrong number of arguments",
			args:               []string{"dummyutil", "deploy", "1.0.0.1"},
			version:            "0.0.0-dummy",
			expectedErr:        nil,
			expectedErrPattern: "5 arg",
		},
		{
			name:               "Wrong argument format",
			args:               []string{"dummyutil", "deploy", "1", "2", "3", "4", "5"},
			version:            "0.0.0-dummy",
			expectedErr:        ErrDeployInvalidArg,
			expectedErrPattern: "invalid argument",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stdout, stderr, err := tests.SilentExecute(execute, tc.args, tc.version)
			fmt.Println(stdout, stderr)

			tests.CheckError(t, tc.expectedErr, tc.expectedErrPattern, err)
		})
	}
}
