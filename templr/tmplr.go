package templr

import (
	"bytes"
	"text/template"
)

func New(opts ...Opt) Templr {
	options := &Opts{}

	var t = new(templr)
	t.opts = options

	configure(t, opts...)

	return t
}

func (t *templr) Parse(s string) string {
	if t.err != nil || s == "" {
		return ""
	}

	templ, err := template.New("").Funcs(templrFuncs).Parse(s)
	if err != nil {
		t.err = err
		return ""
	}

	var b bytes.Buffer
	if err = templ.Execute(&b, t.vars); err != nil {
		t.err = err
		return ""
	}

	return b.String()
}

func (t *templr) Err() error {
	return t.err
}

func configure(t *templr, opts ...Opt) error {
	for _, o := range opts {
		o(t.opts)
	}

	return nil
}
