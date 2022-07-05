package templr

import (
	"bytes"
	"runtime"

	"text/template"

	"golang.org/x/exp/maps"
)

type Templr interface {
	Parse(s string) string
	Err() error
}

type templr struct {
	opts *Opts
	err  error
}

type Vars map[string]Var

type Var string

type Opt func(*Opts)

type Opts struct {
	Fields Fields
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// Fields ...
type Fields map[string]interface{}

// WithVars ...
func WithFields(fields Fields) Opt {
	return func(o *Opts) {
		maps.Copy(o.Fields, fields)
	}
}

// DefaultFields ...
func DefaultFields() (Fields, error) {
	fields := Fields{"OS": runtime.GOOS, "ARCH": runtime.GOARCH}

	return fields, nil
}

// New ...
func New(opts ...Opt) Templr {
	options := &Opts{
		Fields: Fields{"OS": runtime.GOOS, "ARCH": runtime.GOARCH},
	}
	options.Configure(opts...)

	var t = new(templr)
	t.opts = options

	return t
}

// Parse ...
func (t *templr) Parse(s string) string {
	if t.err != nil || s == "" {
		return ""
	}

	tmpl, err := template.New("tmpl").Parse(s)
	if err != nil {
		t.err = err
		return ""
	}

	var b bytes.Buffer
	if err = tmpl.Execute(&b, t.opts.Fields); err != nil {
		t.err = err
		return ""
	}

	return b.String()
}

// Err ...
func (t *templr) Err() error {
	return t.err
}
