package gen

import (
	"os"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

type Generator interface {
	// Write all of the templates
	Write() error
}

type Context interface {
	// Flags is adding flags to the command
	// mapping from the command to the struct
	Flags(cmd *cobra.Command)
	// Execute and write to file
	Execute(t *template.Template, f *os.File) error
}

type generator struct {
	gc   Context
	opts *Opts
	box  *packr.Box
}

type Opt func(*Opts)

type Opts struct {
	// Dir is the directory to write to
	Dir string
}
