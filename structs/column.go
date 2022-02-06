package structs

import (
	"fmt"
	"strings"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/vitess/go/vt/sqlparser"
	"github.com/olachat/gola/mysqldriver"
)

type ColumnStruct struct {
	mysqldriver.Column
	table *TableStruct
}

var dbTypeToSQLTypes = map[string]string{
	"tinyint":           "sql.Int8",
	"smallint":          "sql.Int16",
	"int":               "sql.Int32",
	"bigint":            "sql.Int64",
	"tinyint unsigned":  "sql.Uint8",
	"smallint unsigned": "sql.Uint16",
	"int unsigned":      "sql.Uint32",
	"bigint unsigned":   "sql.Uint64",
}

var dbTypeToGoTypes = map[string]string{
	"tinyint":           "int8",
	"smallint":          "int16",
	"int":               "int",
	"bigint":            "int64",
	"tinyint unsigned":  "uint8",
	"smallint unsigned": "uint16",
	"int unsigned":      "uint",
	"bigint unsigned":   "uint64",
}

var dbTypeToPHPTypes = map[string]string{
	"tinyint":           "int",
	"smallint":          "int",
	"int":               "int",
	"bigint":            "int",
	"tinyint unsigned":  "int",
	"smallint unsigned": "int",
	"int unsigned":      "int",
	"bigint unsigned":   "int",
	"double":            "float",
}

func (c ColumnStruct) SQLType() string {
	if sqlType, ok := dbTypeToSQLTypes[c.FullDBType]; ok {
		return sqlType
	}
	unsignedString := ""
	if strings.HasSuffix(c.FullDBType, "unsigned") {
		unsignedString = " unsigned"
	}
	if strings.HasPrefix(c.DBType, "tinyint") {
		if sqlType, ok := dbTypeToSQLTypes["tinyint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "smallint") {
		if sqlType, ok := dbTypeToSQLTypes["smallint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "bigint") {
		if sqlType, ok := dbTypeToSQLTypes["bigint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "int") {
		if sqlType, ok := dbTypeToSQLTypes["int"+unsignedString]; ok {
			return sqlType
		}
	}

	if strings.HasPrefix(c.DBType, "varchar") {
		size := getValue(c.FullDBType)

		return fmt.Sprintf("sql.MustCreateStringWithDefaults(sqltypes.VarChar, %s)", size)
	}

	if strings.HasPrefix(c.DBType, "decimal") || strings.HasPrefix(c.DBType, "float") {
		return "sql.Float32"
	}

	if strings.HasPrefix(c.DBType, "double") {
		return "sql.Float64"
	}

	if strings.HasPrefix(c.DBType, "enum") {
		enums := strings.ReplaceAll(getValue(c.FullDBType), "'", "\"")

		return fmt.Sprintf("sql.MustCreateEnumType([]string{%s}, sql.Collation_Default)", enums)
	}

	if strings.HasPrefix(c.DBType, "set") {
		vals := strings.ReplaceAll(getValue(c.FullDBType), "'", "\"")

		return fmt.Sprintf("sql.MustCreateSetType([]string{%s}, sql.Collation_Default)", vals)
	}

	if strings.HasPrefix(c.DBType, "char") {
		size := getValue(c.FullDBType)

		return fmt.Sprintf("sql.MustCreateStringWithDefaults(sqltypes.Char, %s)", size)
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "sql.Timestamp"
	}

	if strings.Contains(c.DBType, "text") {
		return "sql.Text"
	}

	if strings.Contains(c.DBType, "blob") {
		return "sql.Blob"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c ColumnStruct) GoType() string {
	if goType, ok := dbTypeToGoTypes[c.DBType]; ok {
		return goType
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "decimal") {
		return "float32"
	}

	if c.IsEnum() {
		return c.table.ClassName() + c.GoName()
	}

	if strings.HasPrefix(c.DBType, "set") {
		return "string"
	}

	if strings.Contains(c.DBType, "text") || strings.HasPrefix(c.DBType, "blob") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "time.Time"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c ColumnStruct) PHPType() string {
	if goType, ok := dbTypeToPHPTypes[c.DBType]; ok {
		return goType
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "decimal") {
		return "float"
	}

	if c.IsEnum() {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "set") {
		return "string"
	}

	if strings.Contains(c.DBType, "text") || strings.HasPrefix(c.DBType, "blob") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "time.Time"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c ColumnStruct) GoName() string {
	return getGoName(c.Name)
}

func (c ColumnStruct) IsNullable() string {
	if c.Nullable {
		return "true"
	}

	return "false"
}

func (c ColumnStruct) HasDefault() bool {
	if c.Column.Default == "auto_increment" {
		return false
	}

	return len(c.Default) > 0
}

func (c ColumnStruct) GoDefaultValue() string {
	goType := c.GoType()
	if goType == "string" || c.IsEnum() {
		if strings.HasPrefix(c.Default, "\"") && strings.HasSuffix(c.Default, "\"") {
			return c.Default
		}
		return "\"" + c.Default + "\""
	}

	if goType == "time.Time" {
		if strings.Contains(c.Default, "CURRENT_TIMESTAMP") {
			return "time.Now()"
		}
		return c.Default
	}

	if strings.Contains(goType, "int") || strings.Contains(goType, "float") {
		return goType + "(" + strings.ReplaceAll(c.Default, `"`, "") + ")"
	}

	return goType + "(" + c.Default + ")"
}

func (c ColumnStruct) GetColumnDefault() string {
	if !c.HasDefault() {
		return ""
	}

	return ", Default: default" + c.GoName()
}

func (c ColumnStruct) IsEnum() bool {
	return strings.HasPrefix(c.DBType, "enum")
}

func (c ColumnStruct) GetEnumConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.GoType() + getGoName(enum) + " " + c.GoType() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

func (c ColumnStruct) IsPrimaryKey() string {
	if c.table.PKey == nil {
		return "false"
	}

	for _, pc := range c.table.PKey.Columns {
		if pc == c.Name {
			return "true"
		}
	}

	return "false"
}

func (c ColumnStruct) IsAutoIncrement() string {
	if c.Default == "auto_increment" {
		return "true"
	}

	return "false"
}
