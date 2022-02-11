package corelib

import (
	"strconv"
	"strings"
)

func JoinInts(vals []int, join string) string {
	strs := make([]string, len(vals))
	for i, id := range vals {
		strs[i] = strconv.Itoa(id)
	}

	return strings.Join(strs, join)
}

func GetParamPlaceHolder(count int) string {
	strs := make([]string, count)
	for i := range strs {
		strs[i] = "?"
	}

	return strings.Join(strs, ",")
}
