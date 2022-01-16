package testdata

import "embed"

//go:embed *.sql */*.go
var Fixtures embed.FS
