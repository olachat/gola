package tests

import (
	"database/sql"
	"fmt"
	"math"
	"testing"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/golalib/testdata"
	"github.com/olachat/gola/v2/golalib/testdata/users"
	"github.com/olachat/gola/v2/mysqldriver"
)

const (
	testDBPort int    = 33067
	testDBName string = "testdb"
)

var tableNames = []string{
	"users", "blogs", "songs", "song_user_favourites", "profile", "account",
	"gifts", "gifts_with_default",
	"gifts_nn", "gifts_nn_with_default", "wallet",
	"worker",
}

func init() {
	engine := sqle.NewDefault(gsql.NewDatabaseProvider(
		memory.NewDatabase(testDBName),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", testDBPort),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()

	connstr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort, "false")
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		panic(err)
	}

	coredb.Setup(func(dbname string, mode coredb.DBMode) *sql.DB {
		if dbname == testDBName {
			return db
		}
		return nil
	})

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test_table (name VARCHAR(255), email VARCHAR(255))")
	if err != nil {
		panic(err)
	}

	// create tables
	for _, tableName := range tableNames {
		query, _ := testdata.Fixtures.ReadFile(tableName + ".sql")
		db.Exec(string(query))
	}

	_, err = db.Exec("SET autocommit = 0")
	if err != nil {
		panic("fail to set autocommit mode")
	}

	// add data
	_, err = db.Exec(`
insert into users (name, email, created_at, updated_at, float_type, double_type, hobby, hobby_no_default, sports_no_default, sports) values
("John Doe", "john@doe.com", NOW(), NOW(), 1.55555, 1.8729, 'running','swimming', ('swim,tennis'), ("tennis")),
("John Doe", "johnalt@doe.com", NOW(), NOW(), 2.5, 2.8239, 'swimming','running', ('basketball'), ("football")),
("Jane Doe", "jane@doe.com", NOW(), NOW(), 3.5, 334.8593, 'singing','swimming', ('squash,badminton'), ("squash,tennis")),
("Evil Bob", "evilbob@gmail.com", NOW(), NOW(), 4.5, 42234.83, 'singing','running', 'tennis', 'badminton,basketball')
	`)
	if err != nil {
		panic("insert failed " + err.Error())
	}

	r, err := db.Exec("INSERT IGNORE INTO `users` (`id`, `name`, `email`, `created_at`, `updated_at`, `float_type`, `double_type`, `hobby`, `hobby_no_default`, `sports`, `sports2`, `sports_no_default`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 40014, "name", "email", uint64(2), uint64(3), 4.5, 5.6, "running", "singing", "swim,tennis", "swim", "")
	if err != nil {
		panic(err.Error())
	}
	if count, err := r.RowsAffected(); err != nil {
		panic(err.Error())
	} else if count == 0 {
		panic("nothing inserted")
	}

	_, err = db.Exec(`
insert into songs (id,title,type,manifest) values
(999,'song1 2 3','101','a')
`)
	if err != nil {
		panic("fail to insert song for testing")
	}

	_, err = db.Exec(`
	insert into gifts (id) values
	(1)
	`)
	if err != nil {
		panic("fail to insert gift. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts (id,name,is_free,gift_count,gift_type,create_time,discount,price,remark,manifest,description,update_time,branches) values
	(2,"name",1,3,'freebie',1678935576,7.5,255.33,'remark is long text','printable manifest','description text','2019-01-01 00:00:01','vivo,sentosa')`)
	if err != nil {
		panic("fail to insert gift. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts_with_default (id) values
	(1)`)
	if err != nil {
		panic("fail to insert gift_with_default. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts_with_default (id,name,is_free,gift_count,gift_type,create_time,discount,price,remark,manifest,description,update_time,branches) values
	(2,"name",1,3,'freebie',1678935576,7.5,255.33,'remark is long text','printable manifest','description text','2019-01-01 00:00:01','vivo,sentosa')`)
	if err != nil {
		panic("fail to insert gift. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts_nn (id,name,is_free,gift_count,gift_type,create_time,discount,price,remark,manifest,description,update_time,branches) values
	(1,"name",1,3,'freebie',1678935576,7.5,255.33,'remark is long text','printable manifest','description text','2019-01-01 00:00:01','vivo,sentosa')`)
	if err != nil {
		panic("fail to insert gift_with_default. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts_nn_with_default (id,description) values
	(1,'gift1')`)
	if err != nil {
		panic("fail to insert gifts_nn_with_default. " + err.Error())
	}

	_, err = db.Exec(`
	insert into gifts_nn_with_default (id,name,is_free,gift_count,gift_type,create_time,discount,price,remark,manifest,description,update_time,branches) values
	(2,"name",1,3,'freebie',1678935576,7.5,255.33,'remark is long text','printable manifest','description text','2019-01-01 00:00:01','vivo,sentosa')
	`)
	if err != nil {
		panic("fail to insert gift. " + err.Error())
	}
}

type SimpleUser struct {
	users.Name
	users.Email
}

func TestUserInsert(t *testing.T) {
	u := users.NewWithPK(11)
	u.SetEmail("hello10023")
	u.SetName("maou sheng")
	u.SetCreatedAt(111)
	u.SetUpdatedAt(222)
	u.SetSports([]users.UserSports{
		users.UserSportsBadminton,
		users.UserSportsBasketball,
	})
	u.SetHobbyNoDefault(users.UserHobbyNoDefaultRunning)
	err := u.Insert()
	if err != nil {
		t.Fatalf("fail to insert: %s", err.Error())
	}
	uOut := users.FetchByPK(11)
	sports := uOut.GetSports()
	if !contains(sports, users.UserSportsBadminton) {
		t.Errorf("uOut.GetSports should contain badminton. Actual: %v", uOut.GetSports())
	}
	if !contains(sports, users.UserSportsBasketball) {
		t.Errorf("uOut.GetSports should contain badminton. Actual: %v", uOut.GetSports())
	}
}

func TestUserDouble(t *testing.T) {
	u1 := users.FetchByPK(1)
	if u1.GetDoubleType() != 1.8729 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u1.GetDoubleType())
	}
	if u1.GetFloatType() != 1.55555 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u1.GetFloatType())
	}

	u2 := users.FetchByPK(2)
	if u2.GetDoubleType() != 2.8239 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u2.GetDoubleType())
	}
	if u2.GetFloatType() != 2.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u2.GetFloatType())
	}

	u3 := users.FetchByPK(3)
	if u3.GetDoubleType() != 334.8593 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u3.GetDoubleType())
	}
	if u3.GetFloatType() != 3.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u3.GetFloatType())
	}

	u4 := users.FetchByPK(4)
	if u4.GetDoubleType() != 42234.83 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u4.GetDoubleType())
	}
	if u4.GetFloatType() != 4.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u4.GetFloatType())
	}

	u4.SetDoubleType(5.1)
	u4.SetFloatType(4.0)
	u4.Update()

	u5 := users.FetchByPK(4)
	if u5.GetDoubleType() != 5.1 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u4.GetDoubleType())
	}
	if u5.GetFloatType() != 4.0 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u4.GetFloatType())
	}
}

func TestUserHobby(t *testing.T) {
	u1 := users.FetchByPK(1)
	if u1.GetHobby() != users.UserHobbyRunning {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u1.GetHobby())
	}
	if u1.GetHobbyNoDefault() != users.UserHobbyNoDefaultSwimming {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u1.GetHobbyNoDefault())
	}

	u1.SetHobby(users.UserHobbySinging)
	u1.Update()
	u1 = users.FetchByPK(1)
	if u1.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u1.GetHobby())
	}

	u2 := users.FetchByPK(2)
	if u2.GetHobby() != users.UserHobbySwimming {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u2.GetHobby())
	}
	if u2.GetHobbyNoDefault() != users.UserHobbyNoDefaultRunning {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u2.GetHobbyNoDefault())
	}

	u3 := users.FetchByPK(3)
	if u3.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u3.GetHobby())
	}
	if u3.GetHobbyNoDefault() != users.UserHobbyNoDefaultSwimming {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u3.GetHobbyNoDefault())
	}

	u4 := users.FetchByPK(4)
	if u4.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u4.GetHobby())
	}
	if u4.GetHobbyNoDefault() != users.UserHobbyNoDefaultRunning {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u4.GetHobbyNoDefault())
	}
}

func TestUserSports(t *testing.T) {
	u1 := users.FetchByPK(1)
	if len(u1.GetSports()) != 1 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u1.GetSports())
	}
	if !contains(u1.GetSports(), users.UserSportsTennis) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u1.GetSports())
	}

	if len(u1.GetSportsNoDefault()) != 2 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u1.GetSportsNoDefault())
	}
	if !contains(u1.GetSportsNoDefault(), users.UserSportsNoDefaultSwim) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain swim. Actual: %v", u1.GetSportsNoDefault())
	}
	if !contains(u1.GetSportsNoDefault(), users.UserSportsNoDefaultTennis) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain tennis. Actual: %v", u1.GetSportsNoDefault())
	}

	u2 := users.FetchByPK(2)
	if len(u2.GetSports()) != 1 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u2.GetSports())
	}
	if !contains(u2.GetSports(), users.UserSportsFootball) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u2.GetSports())
	}

	if len(u2.GetSportsNoDefault()) != 1 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u2.GetSportsNoDefault())
	}
	if !contains(u2.GetSportsNoDefault(), users.UserSportsNoDefaultBasketball) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain basketball. Actual: %v", u2.GetSportsNoDefault())
	}
}

func TestUserSports2(t *testing.T) {
	u3 := users.FetchByPK(3)
	if len(u3.GetSports()) != 2 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u3.GetSports())
	}
	if !contains(u3.GetSports(), users.UserSportsSquash) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u3.GetSports())
	}
	if !contains(u3.GetSports(), users.UserSportsTennis) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u3.GetSports())
	}

	if len(u3.GetSportsNoDefault()) != 2 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u3.GetSportsNoDefault())
	}
	if !contains(u3.GetSportsNoDefault(), users.UserSportsNoDefaultBadminton) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain badminton. Actual: %v", u3.GetSportsNoDefault())
	}
	if !contains(u3.GetSportsNoDefault(), users.UserSportsNoDefaultSquash) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain squash. Actual: %v", u3.GetSportsNoDefault())
	}

	u4 := users.FetchByPK(4)
	if len(u4.GetSports()) != 2 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u4.GetSports())
	}
	if !contains(u4.GetSports(), users.UserSportsBadminton) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u4.GetSports())
	}
	if !contains(u4.GetSports(), users.UserSportsBasketball) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u4.GetSports())
	}

	if len(u4.GetSportsNoDefault()) != 1 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u4.GetSportsNoDefault())
	}
	if !contains(u4.GetSportsNoDefault(), users.UserSportsNoDefaultTennis) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain tennis. Actual: %v", u4.GetSportsNoDefault())
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isFloatSimilar(expected float64, actual float64) bool {
	return math.Abs(expected-actual) < 0.000001
}

func TestUserMethods(t *testing.T) {
	u := users.FetchFieldsByPK[struct {
		users.Email
	}](1)
	if u.GetEmail() != "john@doe.com" {
		t.Error("Failed to FetchByPK with email using id 1")
	}

	u2 := users.FetchFieldsByPK[users.User](1)
	if u2.GetEmail() != "john@doe.com" && u2.GetName() != "John Doe" {
		t.Error("Failed to FetchByPK with User using id 1")
	}
	u2.SetEmail("joe@doe.com")
	u2.SetName("Joe Doe")
	u2.Update()

	u2 = users.FetchFieldsByPK[users.User](1)
	if u2.GetEmail() != "joe@doe.com" && u2.GetName() != "JOe Doe" {
		t.Error("Failed to FetchByPK with User using id 1 after update")
	}
	u2.SetEmail("john@doe.com")
	u2.SetName("John Doe")
	u2.Update()

	u3 := users.FetchByPK(1)
	if u2.GetEmail() != u3.GetEmail() && u2.GetName() != u3.GetName() {
		t.Error("FetchUserByPK and FetchByPK[User] returns different result")
	}

	u4 := users.FetchByPK(0)
	if u4 != nil {
		t.Error("FetchUserByPK must return nil for id 0")
	}

	objs := users.FetchFieldsByPKs[SimpleUser](1, 2)
	if len(objs) != 2 {
		t.Error("FetchByPKs[SimpleUser]([]int{1, 2}) failed")
	}
	if objs[0].GetEmail() != u.GetEmail() {
		t.Error("FetchByPK and FetchByPKs[SimpleUser] returns different result")
	}

	objs2 := users.FetchByPKs(3, 4)
	if len(objs2) != 2 {
		t.Error("FetchUserByPKs([]int{3, 4}) failed")
	}
}
