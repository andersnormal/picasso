package task

import (
	"context"

	"github.com/andersnormal/picasso/executr"
)

func (t *Task) Exec(ctx context.Context) error {
	for _, cmd := range t.Cmds {
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
