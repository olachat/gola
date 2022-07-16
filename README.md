![Builds status](https://github.com/olachat/gola/actions/workflows/go.yml/badge.svg)
![Coverage](badges/coverage.svg)
![Go Report Card](badges/go-report-card.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olachat/gola)
![GitHub repo size](https://img.shields.io/github/repo-size/olachat/gola)
![GitHub issues](https://img.shields.io/github/issues-raw/olachat/gola)
![GitHub pull requests](https://img.shields.io/github/issues-pr/olachat/gola)
![GitHub](https://img.shields.io/github/license/olachat/gola)

# gola

`gola` is a ORM for go utilizing generic with unique design goals.

# Test

`go test` command will:

- Use `testdata` sql to create tables on the fly
- Do code generation for tables
- Compare with `testdata/*.go`
- Report error if file not matching

Use `go test -update`, if template is changed, and want to update `testdata/*.go`

# Todo

- [ ] CURD
  - [x] Insert
    - [x] Default Value
    - [x] LAST_INSERT_ID()
  - [x] Update
    - [ ] Partial Update
    - [ ] Auto updatedt field value
  - [x] Delete
  - [x] Count
  - [ ] Count with IAQ
- [x] Default & per struct connstr
- [ ] Safety
  - [ ] SQL escape
  - [x] parameterize
- [x] code gen template
  - [x] struct
  - [ ] index query methods
    - [x] index
    - [ ] uniuqe index
    - [x] paging & order
      - [ ] paging using cursor
- [ ] better primary key support
  - [x] Single Key types / names
  - [ ] Composite key
- [ ] db types
  - [ ] timestamp
  - [x] boolean
  - [x] float
  - [x] enum
  - [x] set
- [x] Remove sqlboiler dependency in code gen
- [x] Project badges
- [ ] Hooks
  - [ ] Insert
  - [ ] Update
  - [ ] Delete
- [ ] Tests
  - [x] Use sql to create table & insert testdata
  - [ ] Performance test
- [ ] context support
- [ ] transaction support
- [ ] zero reflect verison?
- [ ] Embed groupcache
- [ ] docs & docs & docs...
 - [ ] IAQ: Index Aware Query doc
