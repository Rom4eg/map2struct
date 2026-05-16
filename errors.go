package map2struct

import "errors"

var (
	ErrMustBePointerStruct = errors.New("map2struct: must be pointer struct")
	ErrEmptyMap            = errors.New("map2struct: empty map")
)
