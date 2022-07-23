package coredb

import "errors"

// ErrAvoidInsert represent the error if insertion failed
var ErrAvoidInsert = errors.New("ErrAvoidInsertion")

// ErrAvoidUpdate represent the error if update failed, i.e. no affected row
var ErrAvoidUpdate = errors.New("ErrAvoidUpdate")
