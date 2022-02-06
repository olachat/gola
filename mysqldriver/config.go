package mysqldriver

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Config is a map with helper functions
type Config map[string]interface{}

// MustString retrieves a string that must exist and must be a string, it must also not be the empty string
func (c Config) MustString(key string) string {
	s, ok := c[key]
	if !ok {
		panic(errors.Errorf("failed to find key %s in config", key))
	}

	str, ok := s.(string)
	if !ok {
		panic(errors.Errorf("found key %s in config, but it was not a string (%T)", key, s))
	}

	if len(str) == 0 {
		panic(errors.Errorf("found key %s in config, but it was an empty string", key))
	}

	return str
}

// String retrieves a non-empty string, the bool says if it exists, is of appropriate type,
// and has a non-zero length or not.
func (c Config) String(key string) (string, bool) {
	s, ok := c[key]
	if !ok {
		return "", false
	}

	str, ok := s.(string)
	if !ok {
		return "", false
	}

	if len(str) == 0 {
		return "", false
	}

	return str, true
}

// DefaultString retrieves a non-empty string or the default value provided.
func (c Config) DefaultString(key, def string) string {
	str, ok := c.String(key)
	if !ok {
		return def
	}

	return str
}

// Int retrieves an int, the bool says if it exists, is of the appropriate type,
// and is non-zero. Coerces float64 to int because JSON and Javascript kinda suck.
func (c Config) Int(key string) (int, bool) {
	i, ok := c[key]
	if !ok {
		return 0, false
	}

	var integer int
	switch t := i.(type) {
	case int:
		integer = t
	case float64:
		integer = int(t)
	case string:
		var err error
		integer, err = strconv.Atoi(t)
		if err != nil {
			return 0, false
		}
	default:
		return 0, false
	}

	if integer == 0 {
		return 0, false
	}

	return integer, true
}

// DefaultInt retrieves a non-zero int or the default value provided.
func (c Config) DefaultInt(key string, def int) int {
	i, ok := c.Int(key)
	if !ok {
		return def
	}

	return i
}

// StringSlice retrieves an string slice, the bool says if it exists, is of the appropriate type,
// is non-nil and non-zero length
func (c Config) StringSlice(key string) ([]string, bool) {
	ss, ok := c[key]
	if !ok {
		return nil, false
	}

	var slice []string
	if intfSlice, ok := ss.([]interface{}); ok {
		for _, i := range intfSlice {
			slice = append(slice, i.(string))
		}
	} else if stringSlice, ok := ss.([]string); ok {
		slice = stringSlice
	} else {
		return nil, false
	}

	// Also detects nil
	if len(slice) == 0 {
		return nil, false
	}
	return slice, true
}

// TablesFromList takes a whitelist or blacklist and returns
// the table names.
func TablesFromList(list []string) []string {
	if len(list) == 0 {
		return nil
	}

	var tables []string
	for _, i := range list {
		splits := strings.Split(i, ".")

		if len(splits) == 1 {
			tables = append(tables, splits[0])
		}
	}

	return tables
}

// ColumnsFromList takes a whitelist or blacklist and returns
// the columns for a given table.
func ColumnsFromList(list []string, tablename string) []string {
	if len(list) == 0 {
		return nil
	}

	var columns []string
	for _, i := range list {
		splits := strings.Split(i, ".")

		if len(splits) != 2 {
			continue
		}

		if splits[0] == tablename {
			columns = append(columns, splits[1])
		}
	}

	return columns
}
