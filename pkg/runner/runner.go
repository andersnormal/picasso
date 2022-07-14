package runner

import (
	"context"
	"sync"
)

// Runner ...
type Runner struct {
	ctx   context.Context
	funcs []RunFunc
	pool  sync.Pool
	sync.Mutex
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

	for _, fn := range append(r.funcs, fn...) {
		if err := fn(c); err != nil {
			return err
		}
	}

	return nil
}

// WithContext ...
func WithContext(ctx context.Context) *Runner {
	return &Runner{
		ctx:   ctx,
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
