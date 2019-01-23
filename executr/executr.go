package executr

import (
	"context"
	"os"
	"os/exec"
	"time"
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

	cmd := exec.Command(e.opts.Cmd)
	cmd.Stdin = e.opts.Stdin
	cmd.Stdout = e.opts.Stdout
	cmd.Stderr = e.opts.Stderr

	if e.opts.Timeout == 0 {
		err := cmd.Wait()

		return err
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(e.opts.Timeout)

	select {
	case <-timeout:
		cmd.Process.Kill()
		return ErrTimeout
	case err := <-done:
		return err
	}
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
