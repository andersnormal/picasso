package tmpl

import (
	"bytes"
	"text/template"

	"github.com/andersnormal/picasso/pkg/spec"
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

// ApplyWithPrompts ...
func (t *Template) ApplyWithPrompts(s string, p spec.Placeholders) (string, error) {

	return "", nil
}

// Apply ...
func (t *Template) Apply(s string) (string, error) {
	var out bytes.Buffer

	tmpl, err := template.New("tmpl").Option("missingkey=error").Parse(s)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&out, t.opts.Fields)

	return out.String(), err
}
