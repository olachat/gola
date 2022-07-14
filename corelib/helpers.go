package corelib

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

// GetInterfaceSlice converts []T to []interface{}
func GetInterfaceSlice[T any](data []T) []interface{} {
	result := make([]interface{}, len(data))
	for i, obj := range data {
		result[i] = obj
	}

	return result
}
