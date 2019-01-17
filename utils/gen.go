package utils

import (
	"fmt"
	"os"

	"github.com/gobuffalo/packr/v2"
)

type Generator interface {
	// Write all of the templates
	Write() error
}

type GeneratorContext struct {
	Project string
}

type generator struct {
	gc   *GeneratorContext
	opts *Opts
	box  *packr.Box
}

type Opt func(*Opts)

type Opts struct {
	// Dir is the directory to write to
	Dir string
}

func NewGenerator(b *packr.Box, gc *GeneratorContext, opts ...Opt) Generator {
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

		_, err = fs.WriteString(f.String())
		fmt.Printf("created %s\n", fs.Name())

		fs.Sync()

		return nil
	}
}

func configure(g *generator, opts ...Opt) error {
	for _, o := range opts {
		o(g.opts)
	}

	return nil
}
