package executr

import (
	"errors"
)

var (
	ErrNoCmd   = errors.New("executr: no command to execute")
	ErrTimeout = errors.New("executr: command timeout")
)
