package gen

import (
	"os"

	"github.com/andersnormal/picasso/templr"

	"github.com/gobuffalo/packr/v2"
)

type Generator interface {
	// Write all of the templates
	Write() error
}

type Context interface{}

type generator struct {
	gc        Context
	opts      *Opts
	box       *packr.Box
	templates map[string]string
}

type Opt func(*Opts)

type Opts struct {
	// Dir is the directory to write to
	Dir string
	// FileMode
	FileMode os.FileMode
	// Templates
	Templates map[string]string
	// Vars
	Vars templr.Vars
}
