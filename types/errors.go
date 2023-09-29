package types

import (
	"github.com/pkg/errors"
)

// ErrMalformed is returned when an external representation cannot be turned in to a native representation.
var ErrMalformed = errors.New("malformed representation")
