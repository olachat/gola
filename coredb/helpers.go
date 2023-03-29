package coredb

import (
	"strings"
	"time"
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

func MustParseTime(timestr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05.999999999", timestr)
	if err != nil {
		panic("fail to parse timestr. Error: " + err.Error())
	}
	return t
}

func ValueInSet(set []string, s string) bool {
	for _, v := range set {
		if strings.EqualFold(v, s) {
			return true
		}
	}
	return false
}

func MapSlice[T, K any](data []T, convert func(T) K) []K {
	out := make([]K, len(data))
	for i, obj := range data {
		out[i] = convert(obj)
	}
	return out
}
