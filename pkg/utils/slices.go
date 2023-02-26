/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 */

package coreutils

func Contains[T comparable](elems []T, toFind T) bool {
	for _, elem := range elems {
		if elem == toFind {
			return true
		}
	}
	return false
}
