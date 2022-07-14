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

* Use `testdata` sql to create tables on the fly
* Do code generation for tables
* Compare with `testdata/*.go`
* Report error if file not matching

Use `go test -update`, if template is changed, and want to update `testdata/*.go`

# Todo

* [ ] CURD
  * [X] Insert
    * [X] Default Value
    * [X] LAST_INSERT_ID()
  * [X] Update
    * [ ] Partial Update
    * [ ] Auto updatedt field value
  * [ ] Delete
* [X] Default & per struct connstr
* [ ] Safety
  * [ ] SQL escape
  * [ ] parameterize
* [X] code gen template
  * [X] struct
  * [ ] index query methods
    * [X] index
    * [ ] uniuqe index
    * [X] paging & order
      * [ ] paging using cursor
* [ ] better primary key support
  * [X] Single Key types / names
  * [ ] Composite key
* [ ] db types
  * [ ] timestamp
  * [X] boolean
  * [X] float
  * [X] enum
  * [X] set
* [X] Remove sqlboiler dependency in code gen
* [X] Project badges
* [ ] Tests
  * [X] Use sql to create table & insert testdata
  * [ ] Performance test
* [ ] context support
* [ ] transaction support
* [ ] zero reflect verison?
* [ ] Embed groupcache
* [ ] docs & docs & docs...
