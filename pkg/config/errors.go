package config

import (
	"errors"
)

var (
	ErrNoWatch          = errors.New("picasso: task no defined")
	ErrNoDefaultTask    = errors.New("picasso: default task missing")
	ErrDuplicateDefault = errors.New("picasso: default task duplicate")
)
