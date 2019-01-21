package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/andersnormal/picasso/templates"
	"github.com/gobuffalo/packr/v2"
)

func NewGenerator(b *packr.Box, gc Context, opts ...Opt) Generator {
	options := &Opts{}

	var g = new(generator)
	g.opts = options
	g.gc = gc
	g.box = b
	g.templates = make(templates.Templates)

	configure(g, opts...)

	return g
}

func (g *generator) Write() error {
	err := g.box.Walk(g.writeWalkFunc())
	if err != nil {
		return err
	}

	return nil
}

// +private

func (g *generator) writeWalkFunc() packr.WalkFunc {
	return func(n string, f packr.File) error {
		target := n
		if v, ok := g.templates[n]; ok {
			target = v
		}

		name := filepath.Base(target)
		fp := filepath.Join(g.opts.Dir, filepath.Dir(target))

		if err := MkdirAll(fp, os.FileMode(g.opts.FileMode)); err != nil {
			fmt.Println(err)
			return err
		}

		fs, err := os.Create(filepath.Join(fp, name))
		if err != nil {
			return err
		}
		defer fs.Close()

		t, err := template.New("test").Parse(f.String())
		if err != nil {
			return err
		}

		if err = t.Execute(fs, g.gc); err != nil {
			return err
		}

		return nil
	}
}

func MkdirAll(path string, mode os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, mode)
	}

	return nil
}

func configure(g *generator, opts ...Opt) error {
	for _, o := range opts {
		o(g.opts)
	}

	g.templates = g.opts.Templates

	if g.opts.FileMode == 0 {
		g.opts.FileMode = 0755
	}

	return nil
}
