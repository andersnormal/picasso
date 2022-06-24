package executr

import (
	"context"
	"io"
	"os"
	"strings"

	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

func New(opts ...Opt) Executr {
	options := &Opts{
		Env: make(Env),
	}

	var e = new(exectur)
	e.opts = options

	_ = configure(e, opts...)

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

func (e *exectur) Run(ctx context.Context) error {
	if (e.opts.Cmd) == "" {
		return ErrNoCmd
	}

	cctx, cancel := context.WithTimeout(ctx, e.opts.Timeout)
	defer cancel()

	if e.opts.Timeout != 0 {
		ctx = cctx // use context with timeout
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(e.opts.Cmd), "")
	if err != nil {
		return err
	}

	r, err := interp.New(
		interp.Dir(e.opts.Dir),
		interp.Env(expand.ListEnviron(append(os.Environ(), e.opts.Env.Strings()...)...)),

		interp.Module(interp.DefaultExec),
		interp.Module(interp.OpenDevImpls(interp.DefaultOpen)),

		interp.StdIO(e.opts.Stdin, e.opts.Stdout, e.opts.Stderr),
	)
	if err != nil {
		return err
	}

	err = r.Run(ctx, p)

	// if we just hit the timeout
	if ctx.Err() == context.DeadlineExceeded {
		return ErrTimeout
	}

	return err
}

func configure(e *exectur, opts ...Opt) error {
	for _, o := range opts {
		o(e.opts)
	}

	if e.opts.Stdin == nil {
		e.opts.Stdin = os.Stdin
	}

	if e.opts.Stdout == nil {
		e.opts.Stdout = os.Stdout
	}

	if e.opts.Stderr == nil {
		e.opts.Stderr = os.Stderr
	}

	return nil
}
