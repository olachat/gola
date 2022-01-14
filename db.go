package main

import (
	"fmt"
	"time"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
)

func createTestDatabase() *memory.Database {
	db := memory.NewDatabase(testDBName)
	table := memory.NewTable(tableName, sql.Schema{
		{Name: "id", Type: sql.Int32, Nullable: false, Source: tableName, PrimaryKey: true},
		{Name: "name", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		{Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	})

	db.AddTable(tableName, table)
	ctx := sql.NewEmptyContext()
	table.Insert(ctx, sql.NewRow(1, "John Doe", "john@doe.com", []string{"555-555-555"}, time.Now()))
	table.Insert(ctx, sql.NewRow(2, "John Doe", "johnalt@doe.com", []string{}, time.Now()))
	table.Insert(ctx, sql.NewRow(3, "Jane Doe", "jane@doe.com", []string{}, time.Now()))
	table.Insert(ctx, sql.NewRow(4, "Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
	return db
}

func init() {
	engine := sqle.NewDefault(sql.NewDatabaseProvider(
		createTestDatabase(),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", testDBPort),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()
}
