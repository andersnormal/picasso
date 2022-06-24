package gen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andersnormal/picasso/pkg/templr"
	"github.com/andersnormal/picasso/pkg/utils"

	"github.com/gobuffalo/packr/v2"
)

func NewGenerator(b *packr.Box, gc Context, opts ...Opt) Generator {
	options := &Opts{}

	var g = new(generator)
	g.opts = options
	g.gc = gc
	g.box = b
	g.templates = make(map[string]string)

	_ = configure(g, opts...)

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
		dp := filepath.Join(g.opts.Dir, filepath.Dir(target))

		if err := MkdirAll(dp, os.FileMode(g.opts.FileMode)); err != nil {
			fmt.Println(err)
			return err
		}

		fp := filepath.Join(dp, name)
		if ok, _ := utils.FileExists(fp); !ok {
			fmt.Printf("picasso: %s already exists", fp)
			// do not create, but do not error
			return nil
		}

		fs, err := os.Create(fp)
		if err != nil {
			return err
		}
		defer fs.Close()

		topts := []templr.Opt{
			func(o *templr.Opts) {
				o.Vars = g.opts.Vars
			},
		}

		t := templr.New(topts...)
		_, err = fs.WriteString(t.Parse(f.String()))
		if err != nil {
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
