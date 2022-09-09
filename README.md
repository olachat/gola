![Builds status](https://github.com/olachat/gola/actions/workflows/go.yml/badge.svg)
![Coverage](badges/coverage.svg)
![Go Report Card](badges/go-report-card.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olachat/gola)
![GitHub repo size](https://img.shields.io/github/repo-size/olachat/gola)
![GitHub issues](https://img.shields.io/github/issues-raw/olachat/gola)
![GitHub pull requests](https://img.shields.io/github/issues-pr/olachat/gola)
![GitHub](https://img.shields.io/github/license/olachat/gola)

# gola

`gola` is an ORM lib for go utilizing generic with unique approach:

## Fully strong typed

No need to specify column name via string for all CRUD ops:

```go
b := book.NewBook()
b.SetTitle("The Go Programming Language")
b.SetPrice(10)
b.Insert() // Auto-increment Primary Key id will be assigned

b.SetPrice(20) // Internally mark price field pending for update
b.Update() // Only price field will be updated

book.DeleteByPK(b.GetId())
```


## IAQ - Index Aware Query

Flexible select methods based on db index, i.e. IAQ - Index Aware Query.

Assuming the following table:
```sql
CREATE TABLE `blogs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'User Id',
  `slug` varchar(255) NOT NULL DEFAULT '' COMMENT 'Slug',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'Title',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Category Id',
  `country` varchar(255) NOT NULL DEFAULT '' COMMENT 'Country of the blog user',
  PRIMARY KEY (`id`),
  KEY `user` (`user_id`),
  KEY `country_cate` (`country`, `category_id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

`gola` will generate methods like:

```go
// Select all columns
// data is []*blogs.Blog type
data := blogs.Select().WhereCountryEQ("SG").AndCategoryIdEQ(1).All()

/* Select only id, title columns with order & offset / limit
   data is []*struct {
    blogs.Id
    blogs.Title
  }

  i.e. can only access Id / Title field, fields not selected like country won't be accessible
*/
data := blogs.SelectFields[struct {
    blogs.Id
    blogs.Title
  }]().WhereCountryEQ("SG").AndCategoryIdEQ(1).OrderBy(blogs.IdDesc).Limit(0, 10)
```

The methods like `WhereCountryEQ` and `AndCategoryIdEQ` are **strictly** generated based on index available in db.

If columns are not available in any index, query methods won't be accessible.

In this way, `gola` will then ensure that all queries are backed by index.

PS: Query using arbitrary SQL is also supported, but purposely made not as convenient as IAQ.

## Extreme fast

`gola` is **the fastest** in all benchmarks specified by [boilbench](https://github.com/volatiletech/boilbench/pull/13), comparing to [gorm](https://github.com/go-gorm/gorm) / [xorm](https://gitea.com/xorm/xorm) / [pop](https://github.com/gobuffalo/pop) / [gorp](https://github.com/go-gorp/gorp) / [Kallax](https://github.com/src-d/go-kallax) and [sqlboiler](https://github.com/volatiletech/sqlboiler).

## Misc feature

* JSON Marshal
* Read / Write db separation

# Usage

* Make sure have [go 1.18](https://go.dev/dl/) or above installed
* Install the lastest version of gola binary:

`go install github.com/olachat/gola`

Check gola version: `gola --version`

The lastest version should be `0.1.0`

## Model Generation

After installation, run `gola gen` in a folder containing `gola.toml` config file.

`gola.toml`'s content should be similar to:

```toml
[mysql]
host    = "localhost"
port    = 3306
user    = "root"
pass    = ""
sslmode = "false"
dbname = "testdb"
blacklist = []
whitelist = ["blogs"]
output = "models"
```

`gola gen` will then generate orm codes for `blogs` table into `models` folder.

if whitelist is not specified, all tables in `dbname` database will be generated if not blacklisted.

## Setup

Upon your application start, db instance factory must be passed to gola's `coredb`:

```go
import (
  "database/sql"
  "github.com/olachat/gola/coredb"
)

func main() {
  // dbread := get *sql.DB struct for read
  // dbwrite := get *sql.DB struct for write

  coredb.Setup(func(dbname string, m coredb.DBMode) *sql.DB {
    if m == coredb.DBModeRead {
      return dbread
    }
    return dbwrite
  })
}
```

If you don't need read / write db separation, could just return the same instance:
```go
  // db := get the normal *sql.DB struct

  coredb.Setup(func(_ string, _ coredb.DBMode) *sql.DB {
    return db
  })
```

## CRUD

Common crud operations are all supported.

### Struct Methods

`gola struct` will support:

* `GetXXX` / `SetXXX` methods for all columns fields
* `Insert()`
  * AutoIncrement key value will be automatically assign to struct after insertion
* `Update()`
  * Automatically update fields's `Set` method has been called

### Package Methods

Each table will be generated into its own package / folder.

For example, there are two tables `bloggers` and `blogs` in `blogdb` database, and code generated into `models` folder.

`gola` will create two folders:

* `models/bloggers`
* `models/blogs`

Each folder / package will have following global methods for its table:

* `FetchByPK`: get a row with all column fields from table with given primary key
* `FetchFieldsByPK`: get a row with given columns fields from table with given primary key
* `Select()`: IAQ select methods returns all columns fields
* `SelectFields[T]()`: IAQ select methods returns given columns fields
* `DeleteByPK()`: Delete a row with given primary key

### coredb Methods

* `coredb.Setup`
* `coredb.Exec`
* `coredb.Query[T]`
* `coredb.QueryInt`

# Contribute

Clone the source source:

`git clone git@github.com:olachat/gola.git`

## Test

`go test ./...` will run a full test for gola.

`go test ./golalib` command will:

- Use `golalib/testdata` sql files to create tables on the fly
- Do code generation for tables
- Compare with `golalib/testdata/*.go`
- Report error if file not matching

Use `go test ./golalib -update`, if template is changed, and want to update `golalib/testdata/*.go`

# Todo

- [ ] CURD
  - [ ] Auto updatedt field value
- [ ] index query methods
  - [ ] Count with IAQ
  - [ ] uniuqe index
  - [ ] paging using cursor
- [ ] data types
  - [ ] more nullable types
  - [ ] decimal
- [ ] Hooks
  - [ ] Insert
  - [ ] Update
  - [ ] Delete
- [ ] context support
- [ ] transaction support
- [ ] Cache support
- [ ] tutorial
