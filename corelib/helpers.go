package corelib

import (
	"strconv"
	"strings"
)

// JoinInts returns string of given int slice joined by `join` separator
func JoinInts(vals []int, join string) string {
	strs := make([]string, len(vals))
	for i, id := range vals {
		strs[i] = strconv.Itoa(id)
	}

	return strings.Join(strs, join)
}

// GetParamPlaceHolder returns string for param place holder in sql with given count
func GetParamPlaceHolder(count int) string {
	strs := make([]string, count)
	for i := range strs {
		strs[i] = "?"
	}

	return strings.Join(strs, ",")
}
