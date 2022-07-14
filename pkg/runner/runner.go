package runner

import (
	"context"

	"github.com/andersnormal/picasso/pkg/spec"
)

// Runner ...
type Runner struct {
	ctx   context.Context
	funcs []RunFunc
}

// Context ...
func (r *Runner) Context() context.Context {
	return r.ctx
}

// RunFunc ...
type RunFunc func(c *Ctx) error

// Run ...
func (r *Runner) Run(task *spec.Task, fn RunFunc) error

// WithContext ...
func WithContext(ctx context.Context) *Runner {
	return &Runner{
		ctx,
		make([]RunFunc, 0),
	}
}

// Use ...
func (r *Runner) Use(funcs ...RunFunc) {
	r.funcs = append(r.funcs, funcs...)
}
