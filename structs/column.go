package structs

import (
	"fmt"
	"strings"
)

// modified from: /home/wuvist/go/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/drivers/column.go

// Column holds information about a database column.
// Types are Go types, converted by TranslateColumnType.
type Column struct {
	Name      string `json:"name" toml:"name"`
	Type      string `json:"type" toml:"type"`
	DBType    string `json:"db_type" toml:"db_type"`
	Default   string `json:"default" toml:"default"`
	Comment   string `json:"comment" toml:"comment"`
	Nullable  bool   `json:"nullable" toml:"nullable"`
	Unique    bool   `json:"unique" toml:"unique"`
	Validated bool   `json:"validated" toml:"validated"`

	// Postgres only extension bits
	// ArrType is the underlying data type of the Postgres
	// ARRAY type. See here:
	// https://www.postgresql.org/docs/9.1/static/infoschema-element-types.html
	ArrType *string `json:"arr_type" toml:"arr_type"`
	UDTName string  `json:"udt_name" toml:"udt_name"`
	// DomainName is the domain type name associated to the column. See here:
	// https://www.postgresql.org/docs/10/extend-type-system.html#EXTEND-TYPE-SYSTEM-DOMAINS
	DomainName *string `json:"domain_name" toml:"domain_name"`

	// MySQL only bits
	// Used to get full type, ex:
	// tinyint(1) instead of tinyint
	// Used for "tinyint-as-bool" flag
	FullDBType string `json:"full_db_type" toml:"full_db_type"`

	// MS SQL only bits
	// Used to indicate that the value
	// for this column is auto generated by database on insert (i.e. - timestamp (old) or rowversion (new))
	AutoGenerated bool `json:"auto_generated" toml:"auto_generated"`

	Table *Table
}

var dbTypeToGoTypes = map[string]string{
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int",
	"int":                "int",
	"bigint":             "int64",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint",
	"int unsigned":       "uint",
	"bigint unsigned":    "uint64",
	"float":              "float64",
	"float unsigned":     "float64",
	"double":             "float64",
	"double unsigned":    "float64",
}

func (c Column) GoSetEnumType() string {
	return c.Table.ClassName() + c.GoName()
}
func (c Column) GoSetNullableType() string {
	return fmt.Sprintf("goption.Option[[]%s]", c.Table.ClassName()+c.GoName())
}
func (c Column) GoEnumNullableType() string {
	return fmt.Sprintf("goption.Option[%s]", c.Table.ClassName()+c.GoName())
}

func (c Column) ValType() string {
	if c.FullDBType == "tinyint(1)" {
		if c.Nullable {
			return "goption.Option[int]"
		}
		return "bool"
	}
	if c.IsSet() {
		if c.Nullable {
			return "goption.Option[string]"
		} else {
			return "string"
		}
	}
	return c.GoType()
}

// GoType returns type in go of the column.
// Uses goption for nullable fields
func (c Column) GoType() string {
	if c.FullDBType == "tinyint(1)" || c.FullDBType == "tinyint(1) unsigned" {
		if c.Nullable {
			return "goption.Option[bool]"
		}
		return "bool"
	}

	for dbType, goType := range dbTypeToGoTypes {
		if c.DBType == dbType || strings.HasPrefix(c.DBType, dbType+"(") {
			if c.Nullable {
				return fmt.Sprintf("goption.Option[%s]", goType)
			}
			return goType
		}
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		if c.Nullable {
			return "goption.Option[string]"
		}
		return "string"
	}

	if strings.HasPrefix(c.DBType, "varbinary") || strings.HasPrefix(c.DBType, "binary") {
		if c.Nullable {
			return "goption.Option[[]byte]"
		}
		return "[]byte"
	}

	if strings.HasPrefix(c.DBType, "decimal") {
		if c.Nullable {
			return "goption.Option[float64]"
		}
		return "float64"
	}

	if c.IsEnum() {
		if c.Nullable {
			return fmt.Sprintf("goption.Option[%s]", c.Table.ClassName()+c.GoName())
		}
		return c.Table.ClassName() + c.GoName()
	}

	if c.IsSet() {
		if c.Nullable {
			return fmt.Sprintf("goption.Option[[]%s]", c.Table.ClassName()+c.GoName())
		}
		return c.Table.ClassName() + c.GoName()
	}

	if strings.HasPrefix(c.DBType, "set") {
		if c.Nullable {
			return "goption.Option[string]"
		}
		return "string"
	}

	if strings.Contains(c.DBType, "text") || strings.HasPrefix(c.DBType, "blob") {
		if c.Nullable {
			return "goption.Option[string]"
		}
		return "string"
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		if c.Nullable {
			return "goption.Option[time.Time]"
		}
		return "time.Time"
	}

	panic("Unsupported db type: " + c.DBType)
}

// GoName returns the variable name for go of the column
func (c Column) GoName() string {
	return getGoName(c.Name)
}

// IsNullable returns if the column is nullable as string
func (c Column) IsNullable() bool {
	return c.Nullable
}

// HasDefault returns if the column has default value
func (c Column) HasDefault() bool {
	if c.Default == "auto_increment" {
		return false
	}

	return len(c.Default) > 0
}

func getQuotedStr(str string) string {
	if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
		return str
	}
	return "\"" + str + "\""
}

// GoDefaultValue returns the go value of column's default value
func (c Column) GoDefaultValue() string {
	valType := c.ValType()
	lowerCaseDefault := strings.ToLower(c.Default)

	if c.Nullable && c.Default == "NULL" {
		return strings.ReplaceAll(valType, "Option", "None") + "()"
	}

	if strings.Contains(valType, "goption") && c.IsEnum() {
		return strings.ReplaceAll(valType, "Option", "Some") + "(" + getQuotedStr(c.Default) + ")"
	} else if valType == "string" || c.IsEnum() {
		return getQuotedStr(lowerCaseDefault)
	}

	if valType == "goption.Option[string]" {
		return fmt.Sprintf("goption.Some[string](%s)", getQuotedStr(lowerCaseDefault))
	}
	if valType == "string" || c.IsSet() {
		lowerCaseNoSpaceDefault := strings.ReplaceAll(lowerCaseDefault, " ", "")
		if strings.HasPrefix(lowerCaseNoSpaceDefault, "(") && strings.HasSuffix(lowerCaseNoSpaceDefault, ")") {
			return lowerCaseNoSpaceDefault[1 : len(lowerCaseNoSpaceDefault)-1]
		}
		return getQuotedStr(lowerCaseDefault)
	}

	if valType == "time.Time" {
		if strings.Contains(strings.ToUpper(c.Default), "CURRENT_TIMESTAMP") {
			return "time.Now()"
		}
		return fmt.Sprintf("coredb.MustParseTime(%s)", getQuotedStr(c.Default))
	}
	if valType == "goption.Option[time.Time]" {
		if strings.Contains(strings.ToUpper(c.Default), "CURRENT_TIMESTAMP") {
			return "goption.Some[time.Time](time.Now())"
		}
		return fmt.Sprintf("goption.Some[time.Time](coredb.MustParseTime(%s))", getQuotedStr(c.Default))
	}

	if valType == "bool" {
		if c.Default == "0" {
			return "false"
		}
		return "true"
	} else if valType == "goption.Option[bool]" {
		var s string
		if c.Default == "0" {
			s = "false"
		}
		s = "true"
		return fmt.Sprintf("goption.Some[bool](%s)", s)
	}

	if (strings.Contains(valType, "int") || strings.Contains(valType, "float")) &&
		strings.Contains(valType, "goption") {
		return strings.ReplaceAll(valType, "Option", "Some") + "(" + strings.ReplaceAll(c.Default, `"`, "") + ")"
	} else if strings.Contains(valType, "int") || strings.Contains(valType, "float") {
		return valType + "(" + strings.ReplaceAll(c.Default, `"`, "") + ")"
	}

	if valType == "goption.Option[[]byte]" {
		return strings.ReplaceAll(valType, "Option", "Some") + "([]byte(" + getQuotedStr(c.Default) + "))"
	}

	return valType + "(" + getQuotedStr(c.Default) + ")"
}

// IsEnum returns if column type is enum
func (c Column) IsEnum() bool {
	return strings.HasPrefix(c.DBType, "enum")
}

// IsSet returns if column type is set
func (c Column) IsSet() bool {
	return strings.HasPrefix(c.DBType, "set")
}

// IsBool returns if column type is boolean as tinyint(1)
func (c Column) IsBool() bool {
	return c.FullDBType == "tinyint(1)"
}

// IsNullableBool returns if column type is boolean as tinyint(1) and nullable
func (c Column) IsNullableBool() bool {
	return c.FullDBType == "tinyint(1)" && c.Nullable
}

// GetEnumConst returns enum const definitions in go
func (c Column) GetEnumConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.Table.ClassName() + c.GoName() + getGoName(enum) + " " + c.Table.ClassName() + c.GoName() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

// GetSetConst returns set const definitions in go
func (c Column) GetSetConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.Table.ClassName() + c.GoName() + getGoName(enum) + " " + c.Table.ClassName() + c.GoName() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

// GetSetConstList returns a list of set const definitions in go
func (c Column) GetSetConstList() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = `"` + enum + `"`
	}

	return strings.Join(elements, ",\n") + ","
}

// IsPrimaryKey returns if column is primary key
func (c Column) IsPrimaryKey() bool {
	if c.Table.PKey == nil {
		return false
	}

	for _, pc := range c.Table.PKey.Columns {
		if pc == c.Name {
			return true
		}
	}

	return false
}

// IsAutoIncrement returns if column value is auto incremented
func (c Column) IsAutoIncrement() bool {
	return c.Default == "auto_increment"
}
