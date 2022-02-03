![Builds status](https://github.com/yanghua-ola/gola/actions/workflows/go.yml/badge.svg)
![Coverage](badges/coverage.svg)
![Go Report Card](badges/go-report-card.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olachat/gola)
![GitHub repo size](https://img.shields.io/github/repo-size/olachat/gola) 
![GitHub issues](https://img.shields.io/github/issues-raw/olachat/gola)
![GitHub pull requests](https://img.shields.io/github/issues-pr/olachat/gola)
![GitHub](https://img.shields.io/github/license/yanghua-ola/gola)

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
  * [ ] Insert
  * [ ] Update
  * [ ] Delete
* [ ] Connection Pool
  * [ ] Default & per struct connstr
* [ ] Safety
  * [ ] SQL escape
  * [ ] parameterize
* [X] code gen template
  * [X] struct
  * [ ] index query methods
    * [ ] index
    * [ ] uniuqe index
    * [ ] paging & order
* [ ] better primary key support
  * [ ] Single Key types / names
  * [ ] Composite key
* [ ] db types
  * [ ] timestamp
  * [ ] float
  * [ ] enum
  * [ ] set
* [ ] Remove sqlboiler dependency in code gen
* [ ] Project badges
* [ ] Tests
  * [ ] Use sql to create table & insert testdata
  * [ ] Performance test
* [ ] context support
* [ ] transaction support
* [ ] Non-generice version?
* [ ] zero reflect verison?
* [ ] Embed groupcache
* [ ] docs & docs & docs...
