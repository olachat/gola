package coredb

import "errors"

var ErrAvoidInsert = errors.New("ErrAvoidInsertion")
var ErrAvoidUpdate = errors.New("ErrAvoidUpdate")
var ErrMultipleUpdate = errors.New("ErrMultipleUpdate")
