package task

import (
	"context"
	"os"
	"path"

	"github.com/andersnormal/picasso/pkg/templates"
	"github.com/andersnormal/picasso/pkg/templr"
	"github.com/andersnormal/picasso/pkg/utils"
)

func (t *Task) WriteTemplates(ctx context.Context, cwd string) error {
	for _, tmpl := range t.Templates {
		if err := writeTemplate(ctx, cwd, tmpl); err != nil {
			return err
		}
	}

	return nil
}

func writeTemplate(ctx context.Context, cwd string, tmpl *templates.Template) error {
	in := path.Join(cwd, tmpl.File)
	out := path.Join(cwd, tmpl.Output)

	t, err := utils.Stream(in)
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	topts := []templr.Opt{
		func(o *templr.Opts) {
			o.Vars = make(templr.Vars)
			for k, v := range tmpl.Vars {
				o.Vars[k] = templr.Var(v)
			}
		},
	}

	tr := templr.New(topts...)
	_, err = f.WriteString(tr.Parse(string(t)))
	if err != nil {
		return err
	}

	return nil
}
