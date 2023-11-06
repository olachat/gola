package coredb

import (
	"strings"
)

// GetParamPlaceHolder returns string for param place holder in sql with given count
func GetParamPlaceHolder(count int) string {
	strs := make([]string, count)
	for i := range strs {
		strs[i] = "?"
	}

	return strings.Join(strs, ",")
}

// GetAnySlice converts []T to []any
func GetAnySlice[T any](data []T) []any {
	result := make([]any, len(data))
	for i, obj := range data {
		result[i] = obj
	}

	return result
}
