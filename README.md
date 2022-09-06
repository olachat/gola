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
- [ ] nullable types
- [ ] Hooks
  - [ ] Insert
  - [ ] Update
  - [ ] Delete
- [ ] Tests
  - [ ] Performance test
- [ ] context support
- [ ] transaction support
- [ ] zero reflect verison?
- [ ] Cache support
- [ ] docs & docs & docs...
- [ ] IAQ: Index Aware Query doc
