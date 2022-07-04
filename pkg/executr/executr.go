package executr

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/andersnormal/picasso/pkg/spec"

	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// New ...
func New(opts ...Opt) Executr {
	options := &Opts{
		Env: make(Env),
	}
	options.Configure(opts...)

	e := new(exectur)
	e.opts = options

	return e
}

// Stdin ...
func (e *exectur) Stdin() io.Reader {
	return e.opts.Stdin
}

// Stdout ...
func (e *exectur) Stdout() io.Writer {
	return e.opts.Stdout
}

// Stderr ...
func (e *exectur) Stderr() io.Writer {
	return e.opts.Stderr
}

// Run ...
func (e *exectur) Run(ctx context.Context, task spec.Task) error {
	ctx, cancel := context.WithTimeout(ctx, e.opts.Timeout)
	defer cancel()

	for _, cmd := range task.Commands {
		p, err := syntax.NewParser().Parse(strings.NewReader(string(cmd)), "")
		if err != nil {
			return err
		}

		r, err := interp.New(
			interp.Env(expand.ListEnviron(append(os.Environ(), e.opts.Env.Strings()...)...)),

			interp.Module(interp.DefaultExec),
			interp.Module(interp.OpenDevImpls(interp.DefaultOpen)),

			interp.StdIO(e.opts.Stdin, e.opts.Stdout, e.opts.Stderr),
		)
		if err != nil {
			return err
		}

		err = r.Run(ctx, p)
		if err != nil {
			return err
		}
	}

	return nil
}
