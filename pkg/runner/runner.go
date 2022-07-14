package runner

import (
	"context"
	"io"
	"os"
	"sync"
	"time"
)

// Runner ...
type Runner struct {
	ctx   context.Context
	funcs []RunFunc
	pool  sync.Pool
	sync.Mutex
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
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

// Context ...
func (r *Runner) Context() context.Context {
	return r.ctx
}

// AcquireCtx ...
func (r *Runner) AcquireCtx() *Ctx {
	c := r.pool.Get().(*Ctx)
	c.Reset()

	c.runner = r

	return c
}

// Stdin ...
func (r *Runner) Stdin() io.Reader {
	return r.opts.Stdin
}

// Stdout ...
func (r *Runner) Stdout() io.Writer {
	return r.opts.Stdout
}

// Stderr ...
func (r *Runner) Stderr() io.Writer {
	return r.opts.Stderr
}

// ReleaseFunc ...
type ReleaseFunc func()

// ReleseCtx ...
func (r *Runner) ReleaseCtx(c *Ctx) ReleaseFunc {
	return func() { r.pool.Put(c) }
}

// RunFunc ...
type RunFunc func(c *Ctx) error

// Run ...
func (r *Runner) Run(fn ...RunFunc) error {
	c := r.AcquireCtx()
	defer r.ReleaseCtx(c)

	c.funcs = append(r.funcs, fn...)

	return c.funcs[c.idx](c)
}

// WithContext ...
func WithContext(ctx context.Context, opts ...Opt) *Runner {
	options := new(Opts)
	options.Configure(opts...)

	return &Runner{
		ctx:   ctx,
		opts:  options,
		funcs: make([]RunFunc, 0),
		pool: sync.Pool{
			New: func() interface{} {
				return new(Ctx)
			},
		},
	}
}

// Use ...
func (r *Runner) Use(funcs ...RunFunc) {
	r.funcs = append(r.funcs, funcs...)
}
