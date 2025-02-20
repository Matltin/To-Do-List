package util

import "errors"

var (
	ErrInvalidKey = errors.New("invalid key")
	ErrActivityDone = errors.New("the activity allready marke as done")
	ErrInvalidPageNumber = errors.New("invalid page number")
	ErrInvalidLimitNumber = errors.New("invalid limit number")
)