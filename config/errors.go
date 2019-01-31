package config

import (
	"errors"
)

var (
	ErrNoWatch          = errors.New("picasso: task no defined")
	ErrDuplicateDefault = errors.New("picasso: duplicate default task")
)
