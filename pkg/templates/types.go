package templates

import ()

type Templates []*Template

type Vars map[string]Var

type Var string

type Template struct {
	File        string
	Output      string
	IgnoreError bool `yaml:"ignore_error"`
	Vars        Vars
}
