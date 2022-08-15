package structs

import "strings"

// DBInfo represents information about a database and used for codegen
type DBInfo struct {
	Schema  string   `json:"schema"`
	Tables  []*Table `json:"tables"`
	version string
}

// SetVersion for code gen
func (t *DBInfo) SetVersion(version string) {
	t.version = version
}

// GetVersion for code gen
func (t *DBInfo) GetVersion() string {
	return t.version
}

// GetName for code gen
func (t *DBInfo) GetName() string {
	return t.Schema
}

// Package returns package name
func (t *DBInfo) Package() string {
	return strings.ToLower(t.Schema)
}
