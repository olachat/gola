package mysqldriver

// modified from: https://github.com/volatiletech/sqlboiler/blob/v4.6.0/drivers/keys.go

// PrimaryKey represents a primary key constraint in a database
type PrimaryKey struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
}

// ForeignKey represents a foreign key constraint in a database
type ForeignKey struct {
	Table    string `json:"table"`
	Name     string `json:"name"`
	Column   string `json:"column"`
	Nullable bool   `json:"nullable"`
	Unique   bool   `json:"unique"`

	ForeignTable          string `json:"foreign_table"`
	ForeignColumn         string `json:"foreign_column"`
	ForeignColumnNullable bool   `json:"foreign_column_nullable"`
	ForeignColumnUnique   bool   `json:"foreign_column_unique"`
}
