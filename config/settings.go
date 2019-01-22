package config

import (
	"github.com/andersnormal/picasso/templr"
)

func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) Vars() templr.Vars {
	return templr.Vars{
		"author":  templr.Var(s.Author),
		"project": templr.Var(s.Project),
	}
}
