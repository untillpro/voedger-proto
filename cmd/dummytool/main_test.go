package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicUsage_Deploy(t *testing.T) {

	require := require.New(t)

	args := []string{"deploy"}
	err := execute(args, version)
	require.Error(err)
	require.Contains(err.Error(), "5 arg")
}
