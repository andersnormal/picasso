package plugin

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/andersnormal/picasso/pkg/specs"
)

// Run ...
func (o Options) Run(f func(*Plugin) error) {
	if err := run(o, f); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

func run(opts Options, f func(*Plugin) error) error {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	s := bytes.NewReader(in)

	var spec specs.Tmpl
	dec := gob.NewDecoder(s)
	err = dec.Decode(&spec)
	if err != nil {
		return err
	}

	gen, err := opts.New()
	if err != nil {
		return err
	}

	if err := f(gen); err != nil {
		gen.Error(err)
	}

	return nil
}

// Options ...
type Options struct {
	// ParamFunc ...
	paramFunc func(name, value string) error
}

// Plugin ...
type Plugin struct {
	Spec specs.Tmpl

	opts Options
	err  error
}

// New returns a new Plugin.
func (opts Options) New() (*Plugin, error) {
	gen := &Plugin{}

	return gen, nil
}

// Error ...
func (gen *Plugin) Error(err error) {
	if gen.err == nil {
		gen.err = err
	}
}
