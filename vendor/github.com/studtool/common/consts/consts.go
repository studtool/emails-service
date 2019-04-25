package consts

import (
	"time"
)

const (
	EmptyString = ""

	AnyInt    = 0
	AnyBool   = false
	AnyString = EmptyString
	AnyTime   = time.Duration(AnyInt)
)
