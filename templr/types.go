package templr

type Templr interface {
	Parse(s string) string
	Err() error
}

type templr struct {
	vars Vars

	opts *Opts
	err  error
}

type Vars map[string]Var

type Var string

type Opt func(*Opts)

type Opts struct {
}

// todo: move to templates
// type Templates map[string]*Template

// type Template struct {
// 	File        string
// 	Output      string
// 	IgnoreError bool `yaml:"ignore_error"`
// 	Vars        Vars
// }
