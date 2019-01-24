package executr

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
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
		ctx = cctx // use context withtimeout
	}

	f := strings.Fields(e.opts.Cmd)
	name := f[0]
	args := []string{}
	if len(f) > 1 {
		args = append(args, f[1:]...)
	}
	cmd := exec.CommandContext(ctx, name, args...)

	// set env
	cmd.Env = append(os.Environ(), e.opts.Env.Strings()...)

	// setting output
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
