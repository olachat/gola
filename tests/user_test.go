package user

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/mysqldriver"
	"github.com/olachat/gola/testdata"
	"github.com/olachat/gola/testdata/users"
)

const (
	testDBPort int    = 33067
	testDBName string = "testdb"
	tableName  string = "users"
)

func init() {
	corelib.Setup(fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))

	engine := sqle.NewDefault(gsql.NewDatabaseProvider(
		memory.NewDatabase(testDBName),
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

	connStr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort, "false")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	//create table
	query, _ := testdata.Fixtures.ReadFile(tableName + ".sql")
	db.Exec(string(query))

	//add data
	db.Exec(`
insert into users (name, email, created_at, updated_at) values
("John Doe", "john@doe.com", NOW(), NOW()),
("John Doe", "johnalt@doe.com", NOW(), NOW()),
("Jane Doe", "jane@doe.com", NOW(), NOW()),
("Evil Bob", "evilbob@gmail.com", NOW(), NOW())
	`)
}

type SimpleUser struct {
	users.Name
	users.Email
}

func TestUserMethods(t *testing.T) {

	ctx := context.Background()

	u := users.FetchById[struct {
		users.Email
	}](ctx, 1)
	if u.GetEmail(ctx) != "john@doe.com" {
		t.Error("Failed to FetchById with email using id 1")
	}

	u2 := users.FetchById[users.User](ctx, 1)
	if u2.GetEmail(ctx) != "john@doe.com" && u2.GetName(ctx) != "John Doe" {
		t.Error("Failed to FetchById with User using id 1")
	}

	u3 := users.FetchUserById(ctx, 1)
	if u2.GetEmail(ctx) != u3.GetEmail(ctx) && u2.GetName(ctx) != u3.GetName(ctx) {
		t.Error("FetchUserById and FetchById[User] returns different result")
	}

	u4 := users.FetchUserById(ctx, 0)
	if u4 != nil {
		t.Error("FetchUserById must return nil for id 0")
	}

	objs := users.FetchByIds[SimpleUser](ctx, []int{1, 2})
	if len(objs) != 2 {
		t.Error("FetchByIds[SimpleUser]([]int{1, 2}) failed")
	}
	if objs[0].GetEmail(ctx) != u.GetEmail(ctx) {
		t.Error("FetchById and FetchByIds[SimpleUser] returns different result")
	}

	objs2 := users.FetchUserByIds(ctx, []int{3, 4})
	if len(objs2) != 2 {
		t.Error("FetchUserByIds([]int{3, 4}) failed")
	}
}
