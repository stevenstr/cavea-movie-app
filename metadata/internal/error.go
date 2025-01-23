package internal

import "errors"

// ErrNotFound is returned when a requested record
// was not found.
var ErrNotFound = errors.New("not found")
