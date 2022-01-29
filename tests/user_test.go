package user

import (
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
insert into users (name, email, created_at, updated_at, float_type, double_type, hobby) values
("John Doe", "john@doe.com", NOW(), NOW(), 1.55555, 1.8729, 'running'),
("John Doe", "johnalt@doe.com", NOW(), NOW(), 2.5, 2.8239, 'swimming'),
("Jane Doe", "jane@doe.com", NOW(), NOW(), 3.5, 334.8593, 'singing'),
("Evil Bob", "evilbob@gmail.com", NOW(), NOW(), 4.5, 42234.83, 'singing')
	`)
}

type SimpleUser struct {
	users.Name
	users.Email
}

func TestUserDouble(t *testing.T) {
	u1 := users.FetchUserById(1)
	if u1.GetDoubleType() != 1.8729 {
		t.Errorf("FetchUserById GetDoubleType returns unexpected value: %f", u1.GetDoubleType())
	}
	if u1.GetFloatType() != 1.55555 {
		t.Errorf("FetchUserById GetFloatType returns unexpected value: %f", u1.GetFloatType())
	}

	u2 := users.FetchUserById(2)
	if u2.GetDoubleType() != 2.8239 {
		t.Errorf("FetchUserById GetDoubleType returns unexpected value: %f", u2.GetDoubleType())
	}
	if u2.GetFloatType() != 2.5 {
		t.Errorf("FetchUserById GetFloatType returns unexpected value: %f", u2.GetFloatType())
	}

	u3 := users.FetchUserById(3)
	if u3.GetDoubleType() != 334.8593 {
		t.Errorf("FetchUserById GetDoubleType returns unexpected value: %f", u3.GetDoubleType())
	}
	if u3.GetFloatType() != 3.5 {
		t.Errorf("FetchUserById GetFloatType returns unexpected value: %f", u3.GetFloatType())
	}

	u4 := users.FetchUserById(4)
	if u4.GetDoubleType() != 42234.83 {
		t.Errorf("FetchUserById GetDoubleType returns unexpected value: %f", u4.GetDoubleType())
	}
	if u4.GetFloatType() != 4.5 {
		t.Errorf("FetchUserById GetFloatType returns unexpected value: %f", u4.GetFloatType())
	}
}

func TestUserHobby(t *testing.T) {
	u1 := users.FetchUserById(1)
	if u1.GetHobby() != users.UserHobbyRunning {
		t.Errorf("FetchUserById GetHobby returns unexpected value: %v", u1.GetHobby())
	}

	u2 := users.FetchUserById(2)
	if u2.GetHobby() != users.UserHobbySwimming {
		t.Errorf("FetchUserById GetHobby returns unexpected value: %v", u2.GetHobby())
	}

	u3 := users.FetchUserById(3)
	if u3.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserById GetHobby returns unexpected value: %v", u3.GetHobby())
	}

	u4 := users.FetchUserById(4)
	if u4.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserById GetHobby returns unexpected value: %v", u4.GetHobby())
	}
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
