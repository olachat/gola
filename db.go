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
	"github.com/olachat/gola/corelib"
)

func createTestDatabase() *memory.Database {
	db := memory.NewDatabase(corelib.TestDBName)
	table := memory.NewTable(corelib.TableName, sql.Schema{
		{Name: "id", Type: sql.Int32, Nullable: false, Source: corelib.TableName, PrimaryKey: true},
		{Name: "name", Type: sql.Text, Nullable: false, Source: corelib.TableName},
		{Name: "email", Type: sql.Text, Nullable: false, Source: corelib.TableName},
		{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: corelib.TableName},
		{Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: corelib.TableName},
	})

	db.AddTable(corelib.TableName, table)
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
		Address:  fmt.Sprintf("localhost:%d", corelib.TestDBPort),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()
}
