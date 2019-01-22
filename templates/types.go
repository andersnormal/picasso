package templates

import (
	"github.com/andersnormal/picasso/templr"
)

type Templates []*Template

type Template struct {
	File        string
	Output      string
	IgnoreError bool `yaml:"ignore_error"`
	Vars        templr.Vars
}
