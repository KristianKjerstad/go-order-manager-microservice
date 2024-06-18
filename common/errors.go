package common

import "errors"

var (
	ErrorNoItems = errors.New("Items must have at least one item.")
)
