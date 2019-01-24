package cmd

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/andersnormal/picasso/config"
	"github.com/andersnormal/picasso/executr"
	s "github.com/andersnormal/picasso/settings"

	"github.com/spf13/cobra"
)

func init() {}

var (
	hooks = []string{"preRun", "postRun"}
)

var Build = &cobra.Command{
	Use:   "build",
	Short: "build a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
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
		ss.Read(&settings)

		for _, hook := range hooks {
			task, ok := settings.Build[hook]
			if !ok {
				continue
			}

			fmt.Printf("executing %s", task.Desc)

			err := execTask(ctx, task)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func execTask(ctx context.Context, task *config.Task) error {
	for _, cmd := range task.Cmds {
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
