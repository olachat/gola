package main

import (
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/olachat/gola/corelib"
)

const (
	TestDBPort int    = 33066
	TestDBName string = "testdb"
	TableName  string = "users"
)

func createTestDatabase() *memory.Database {
	db := memory.NewDatabase(TestDBName)
	return db
}

type User struct {
	Name
}

type Name string

func (c Name) GetName() string {
	return string(c)
}

func (c *Name) SetName(val string) {
	*c = Name(val)
}

func (c Name) GetValPointer() interface{} {
	return &c
}

func main() {
	var c Name = "hello"
	fmt.Printf("c: %v\n", c)
	c.SetName("world")
	fmt.Printf("c: %v\n", c)

	u := new(User)
	u.Name = "foo"
	fmt.Printf("u: %v\n", u)
	u.SetName("bar")
	fmt.Printf("u: %v\n", u)
}

func main1() {

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

	s.Start()

}
