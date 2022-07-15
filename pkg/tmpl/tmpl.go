package tmpl

import (
	"bytes"
	"text/template"
)

// Template ...
type Template struct {
	opts TmplOpts
}

// Fields ...
type Fields map[string]interface{}

// New ...
func New(opts ...TmplOpt) *Template {
	options := NewOpts()
	options.Configure(opts...)

	return &Template{
		opts: options,
	}
}

// Apply ...
func (t *Template) Apply(s string) (string, error) {

	tmpl, err := template.New("").Funcs(t.opts.Funcs).Parse(s)
	if err != nil {
		return "", err
	}

	if t.opts.FailOnMissing {
		tmpl.Option("missingkey=error")
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, t.opts.Fields); err != nil {
		return "", err
	}

	if t.opts.DisableReplaceNoValue {
		return out.String(), nil
	}

	b := bytes.ReplaceAll(out.Bytes(), []byte("<no value>"), []byte(""))

	return string(b), err
}
