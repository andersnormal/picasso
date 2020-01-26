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
	Vars Vars
}
