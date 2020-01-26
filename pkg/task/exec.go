package task

import (
	"context"

	"github.com/andersnormal/picasso/pkg/executr"
)

func (t *Task) Exec(ctx context.Context) error {
	return execCmds(ctx, t.Cmds)
}

func (t *Task) ExecDeps(ctx context.Context) error {
	for _, dep := range t.resolvedDeps {
		if err := execCmds(ctx, dep.Cmds); err != nil {
			return err
		}
	}

	return nil
}

func execCmds(ctx context.Context, cmds []Cmd) error {
	for _, cmd := range cmds {
		opts := []executr.Opt{func(o *executr.Opts) {
			o.Cmd = string(cmd)
		}}

		e := executr.New(opts...)

		if err := e.Run(ctx); err != nil {
			return err
		}
	}

	return nil
}
