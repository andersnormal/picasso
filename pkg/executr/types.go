package executr

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"
)

// Executr ...
type Executr interface {
	// Run ...
	Run(ctx context.Context, task spec.Task) error
	// Stdin ...
	Stdin() io.Reader
	// Stdout ...
	Stdout() io.Writer
	// Stderr ...
	Stderr() io.Writer
}

// Cmd ...
type Cmd string

// Env ...
type Env map[string]string

// Strings ...
func (ev Env) Strings() []string {
	var env []string
	for k, v := range ev {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env
}

type exectur struct {
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Cmd     string
	Dir     string
	Env     Env
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Timeout time.Duration
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}

	if o.Stdin == nil {
		o.Stdin = os.Stdin
	}

	if o.Stdout == nil {
		o.Stdout = os.Stdout
	}

	if o.Stderr == nil {
		o.Stderr = os.Stderr
	}
}

// WithEnv ...
func WithEnv(env Env) Opt {
	return func(o *Opts) {
		o.Env = env
	}
}

// WithTimeout ...
func WithTimeout(timeout time.Duration) Opt {
	return func(o *Opts) {
		o.Timeout = timeout
	}
}
