// Code generated by gola 0.0.3; DO NOT EDIT.

package testdata

import (
	"database/sql"
	"github.com/olachat/gola/coredb"
)

// Setup default db conn for `testdata` database
// If not set, it will fallback to gola's default db conn
func Setup(db *sql.DB) {
	coredb.SetupDB("testdata", db)
}
