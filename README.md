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

b.SetPrice(20)
b.Update() // Only price field will be updated

book.DeleteByPK(b.GetId())
```


## IAQ - Index Aware Query

Flexible select methods based on db index, i.e. IAQ - Index Aware Query.

## Extreme fast

## Misc

* JSON Marshal

# Usage

* Make use have [go 1.18](https://go.dev/dl/) or above installed
* Install the lastest version of gola binary:

`go install github.com/olachat/gola`

Check gola version: `gola --version`

The lastest version should be `0.1.0`

## Model Generation




## Setup


## CRUD

### Struct Methods

Insert

AutoIncrement Key

Update

  Partial

### Package Methods

FetchByPK
FetchFieldsByPK
SelectFields
DeleteByPK


### coredb Methods

```
coredb.Setup
coredb.Exec
coredb.Query[T]
coredb.QueryInt
```

# Contribute

Clone the source source:

`git clone git@github.com:olachat/gola.git`

## Test

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
  - [ ] nullable types
  - [ ] decimal
- [ ] Hooks
  - [ ] Insert
  - [ ] Update
  - [ ] Delete
- [ ] context support
- [ ] transaction support
- [ ] zero reflect verison?
- [ ] Cache support
- [ ] docs & docs & docs...
- [ ] IAQ: Index Aware Query doc
