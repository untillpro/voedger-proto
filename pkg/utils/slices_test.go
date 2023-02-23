/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 */

package coreutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	require := require.New(t)
	t.Run("basic", func(t *testing.T) {
		bb := []byte{1, 2, 3}
		require.True(Contains(bb, byte(1)))
		require.False(Contains(bb, byte(4)))
	})

	t.Run("structs", func(t *testing.T) {
		type myType struct {
			i int
			s string
		}
		x := myType{
			i: 42,
			s: "str",
		}
		xes := []myType{x, {}}
		require.True(Contains(xes, x))
		require.True(Contains(xes, myType{i: 42, s: "str"}))
		require.False(Contains(xes, myType{i: 42, s: "str1"}))
		require.True(Contains(xes, myType{}))
	})
}
