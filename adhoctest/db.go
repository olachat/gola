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

type idxQuery struct {
}

func (c *idxQuery) WhereNameEQ() idxQuery2 {
	return c
}

func (c *idxQuery) AndSlugEQ() idxQuery3 {
	return c
}

func (c *idxQuery) AndIsVipEQ() idxQuery4 {
	return &idx2{c}
}

type idx2 struct {
	*idxQuery
}

func (c *idx2) AndSlugEQ() idxQuery1 {
	return c
}

type idxQuery1 interface {
	WhereNameEQ() idxQuery2
}

type idxQuery2 interface {
	AndSlugEQ() idxQuery3
}

type idxQuery3 interface {
	AndIsVipEQ() idxQuery4
}

type idxQuery4 interface {
	AndSlugEQ() idxQuery1
}

func Select() idxQuery1 {
	return new(idxQuery)
}

func main() {
	s := Select()
	fmt.Printf("s: %v\n", s)
	a := s.WhereNameEQ()
	fmt.Printf("a: %v\n", a)
	b := a.AndSlugEQ()
	fmt.Printf("b: %v\n", b)
	c := b.AndIsVipEQ()
	fmt.Printf("c: %v\n", c)
	d := c.AndSlugEQ()
	fmt.Printf("d: %v\n", d)
	e := d.WhereNameEQ()
	fmt.Printf("e: %v\n", e)
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
