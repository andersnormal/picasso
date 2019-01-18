package gen

import (
	"os"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

func NewGenerator(b *packr.Box, gc Context, opts ...Opt) Generator {
	options := &Opts{}

	var g = new(generator)
	g.opts = options
	g.gc = gc
	g.box = b

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
	return func(s string, f packr.File) error {
		fs, err := os.Create(s)
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

func configure(g *generator, opts ...Opt) error {
	for _, o := range opts {
		o(g.opts)
	}

	return nil
}
