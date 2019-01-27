package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/andersnormal/picasso/config"
	s "github.com/andersnormal/picasso/settings"
	"github.com/andersnormal/picasso/watcher"

	"github.com/spf13/cobra"
)

func init() {
}

var Watch = &cobra.Command{
	Use:   "watch",
	Short: "watch your developmetn",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("watching files (press ctrl-c to stop) ...")

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
		task, err := settings.Task("watch")
		if err != nil {
			return err
		}

		// watcher opts
		wopts := []watcher.Opt{func(o *watcher.Opts) {
			o.Paths = task.Paths
			o.Cwd = cwd
		}}

		// create watcher
		w := watcher.New(task, wopts...)
		err = w.Reload(ctx)

		return nil
	},
}
