package user

import (
	"fmt"
	"testing"
	"time"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/testdata/users"
)

const (
	TestDBPort int    = 33067
	TestDBName string = "testdb"
	TableName  string = "users"
)

func createTestDatabase() *memory.Database {
	db := memory.NewDatabase(TestDBName)
	table := memory.NewTable(TableName, sql.Schema{
		{Name: "id", Type: sql.Int32, Nullable: false, Source: TableName, PrimaryKey: true},
		{Name: "name", Type: sql.Text, Nullable: false, Source: TableName},
		{Name: "email", Type: sql.Text, Nullable: false, Source: TableName},
		{Name: "created_at", Type: sql.Uint32, Nullable: false, Source: TableName},
		{Name: "updated_at", Type: sql.Uint32, Nullable: false, Source: TableName},
	})

	db.AddTable(TableName, table)
	ctx := sql.NewEmptyContext()
	table.Insert(ctx, sql.NewRow(1, "John Doe", "john@doe.com", time.Now(), time.Now()))
	table.Insert(ctx, sql.NewRow(2, "John Doe", "johnalt@doe.com", time.Now(), time.Now()))
	table.Insert(ctx, sql.NewRow(3, "Jane Doe", "jane@doe.com", time.Now(), time.Now()))
	table.Insert(ctx, sql.NewRow(4, "Evil Bob", "evilbob@gmail.com", time.Now(), time.Now()))
	return db
}

func init() {
	corelib.Setup(fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", TestDBPort, TestDBName))

	engine := sqle.NewDefault(sql.NewDatabaseProvider(
		createTestDatabase(),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", TestDBPort),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()
}

type SimpleUser struct {
	users.Name
	users.Email
}

func TestUserMethods(t *testing.T) {
	u := users.FetchById[struct {
		users.Email
	}](1)
	if u.GetEmail() != "john@doe.com" {
		t.Error("Failed to FetchById with email using id 1")
	}

	u2 := users.FetchById[users.User](1)
	if u2.GetEmail() != "john@doe.com" && u2.GetName() != "John Doe" {
		t.Error("Failed to FetchById with User using id 1")
	}

	u3 := users.FetchUserById(1)
	if u2.GetEmail() != u3.GetEmail() && u2.GetName() != u3.GetName() {
		t.Error("FetchUserById and FetchById[User] returns different result")
	}

	u4 := users.FetchUserById(0)
	if u4 != nil {
		t.Error("FetchUserById must return nil for id 0")
	}

	objs := users.FetchByIds[SimpleUser]([]int{1, 2})
	if len(objs) != 2 {
		t.Error("FetchByIds[SimpleUser]([]int{1, 2}) failed")
	}
	if objs[0].GetEmail() != u.GetEmail() {
		t.Error("FetchById and FetchByIds[SimpleUser] returns different result")
	}

	objs2 := users.FetchUserByIds([]int{3, 4})
	if len(objs2) != 2 {
		t.Error("FetchUserByIds([]int{3, 4}) failed")
	}
}
