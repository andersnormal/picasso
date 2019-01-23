package executr

import (
	"context"
	"os"
	"os/exec"
)

func New(opts ...Opt) Executr {
	options := &Opts{
		Env: make(Env),
	}

	var e = new(exectur)
	e.opts = options

	configure(e, opts...)

	return e
}

func (e *exectur) Run(ctx context.Context) error {
	if (e.opts.Cmd) == "" {
		return ErrNoCmd
	}

	cctx, cancel := context.WithTimeout(ctx, e.opts.Timeout)
	defer cancel()

	if e.opts.Timeout != 0 {
		ctx = cctx // use context withtimeout
	}

	cmd := exec.CommandContext(ctx, e.opts.Cmd)
	cmd.Stdin = e.opts.Stdin
	cmd.Stdout = e.opts.Stdout
	cmd.Stderr = e.opts.Stderr

	err := cmd.Run() // wait for the result of the command

	// if we just hit the timeout
	if ctx.Err() == context.DeadlineExceeded {
		return ErrTimeout
	}

	return err
}

func (e *exectur) Env() []string {
	environ := os.Environ()
	environ = append(environ, e.opts.Env.Strings()...)

	return environ
}

func configure(e *exectur, opts ...Opt) error {
	for _, o := range opts {
		o(e.opts)
	}

	return nil
}
