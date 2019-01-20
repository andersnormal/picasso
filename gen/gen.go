package gen

import (
	"os"
	"path/filepath"
	"strings"
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
	return func(n string, f packr.File) error {
		name := strings.TrimPrefix(filepath.Base(n), "_")
		fp := filepath.Join(g.opts.Dir, filepath.Dir(n))

		if err := MkdirAll(fp, g.opts.FileMode); err != nil {
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

	if g.opts.FileMode == 0 {
		g.opts.FileMode = 0777
	}

	return nil
}
