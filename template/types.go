package template

type Templates map[string]*Template

type Template struct {
	File        string
	Output      string
	IgnoreError bool `yaml:"ignore_error"`
	Vars        Vars
}

type Vars map[string]Var

type Var interface{}
