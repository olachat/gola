package mysqlparser

import (
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func getColComments(opts []*sqlparser.ColumnOption) string {
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionComment {
			return opt.Value
		}
	}
	return ""
}

func getColDefault(opts []*sqlparser.ColumnOption) string {
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionDefaultValue {
			return strings.ReplaceAll(opt.Value, "\"", "")
		}
	}
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionAutoIncrement {
			return "auto_increment"
		}
	}
	return ""
}

func isColNullable(opts []*sqlparser.ColumnOption) bool {
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionPrimaryKey {
			return false
		}
		if opt.Type == sqlparser.ColumnOptionNull {
			return true
		}
		if opt.Type == sqlparser.ColumnOptionNotNull {
			return false
		}
	}
	return true
}

func isColPrimaryKey(opts []*sqlparser.ColumnOption) bool {
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionPrimaryKey {
			return true
		}
	}
	return false
}

func isColUnique(opts []*sqlparser.ColumnOption) bool {
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionPrimaryKey {
			return true
		}
	}
	for _, opt := range opts {
		if opt.Type == sqlparser.ColumnOptionUniqKey {
			return true
		}
	}
	return false
}
