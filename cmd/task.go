package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/andersnormal/picasso/pkg/config"
	s "github.com/andersnormal/picasso/pkg/settings"
	"github.com/andersnormal/picasso/pkg/task"
	"github.com/andersnormal/picasso/pkg/watcher"

	"github.com/spf13/cobra"
)

func init() {
}

func generateTask(name string, task *task.Task) *cobra.Command {
	use := name

	return &cobra.Command{
		Use:   name,
		Short: task.Desc,
		RunE: func(cmd *cobra.Command, args []string) error {
			// create cancable context
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// configure path
			cwd, err := cfg.Cwd()
			if err != nil {
				return err
			}

			// settings opts
			sopts := []s.Opt{func(o *s.Opts) {
				o.File = path.Join(cwd, cfg.File)
				o.FileMode = cfg.FileMode
			}}

			// new settings
			settings := config.NewSettings()
			ss := s.New(sopts...)
			if err := ss.Read(&settings); err != nil {
				return err
			}

			// get task
			task, err := settings.Task(use)
			if err != nil {
				return err
			}

			// write templates
			if err := task.WriteTemplates(ctx, cwd); err != nil {
				return err
			}

			// first execute deps
			if err := task.ExecDeps(ctx); err != nil {
				return err
			}

			// then execute the cmds
			if err := task.Exec(ctx); err != nil {
				return err
			}

			if task.ShouldWatch() {
				fmt.Println("watching files (press ctrl-c to stop) ...")

				// watcher opts
				wopts := []watcher.Opt{func(o *watcher.Opts) {
					o.Paths = task.Paths
					o.Cwd = cwd
				}}

				// create watcher
				w := watcher.New(task, wopts...)
				err = w.Reload(ctx)
			}

			return err
		},
	}
}
