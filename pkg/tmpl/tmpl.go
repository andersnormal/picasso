package tmpl

import (
	"bytes"
	"text/template"

	"github.com/andersnormal/picasso/pkg/spec"

	"github.com/manifoldco/promptui"
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
	var out bytes.Buffer

	tmpl, err := template.New("tmpl").Option("missingkey=error").Parse(s)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&out, t.opts.Fields)

	return out.String(), err
}

// ApplyPrompts ...
func (t *Template) ApplyPrompts(pp spec.Inputs) error {
	ff := make(Fields)

	for _, p := range pp {
		prompt := promptui.Prompt{
			Label: p.Prompt,
		}

		res, err := prompt.Run()
		if err != nil {
			return err
		}
		ff[p.Name] = res
	}

	for f, v := range ff {
		t.opts.Fields[f] = v
	}

	return nil
}
