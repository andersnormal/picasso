package templr

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"

	"text/template"

	"golang.org/x/exp/maps"
)

type Templr interface {
	Parse(s string) string
	ParseFile(string, string) error
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
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fields := Fields{
		"OS":   runtime.GOOS,
		"ARCH": runtime.GOARCH,
		"CWD":  cwd,
	}

	return fields, nil
}

// New ...
func New(opts ...Opt) Templr {
	options := &Opts{
		Fields: make(Fields),
	}
	options.Configure(opts...)

	var t = new(templr)
	t.opts = options

	return t
}

// ParseFile ...
func (t *templr) ParseFile(in, out string) error {
	i, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}

	o, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer o.Close()

	tmpl, err := template.New("tmpl").Parse(string(i))
	if err != nil {
		return err
	}

	err = tmpl.Execute(o, t.opts.Fields)
	if err != nil {
		return err
	}

	return nil
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
