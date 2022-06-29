package cmd

import (
	"context"

	"github.com/andersnormal/picasso/pkg/gen"
	"github.com/andersnormal/picasso/pkg/gen/iface"
	"github.com/andersnormal/picasso/pkg/plugin"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                "init",
	Short:              "Initialized a new project",
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := zap.NewProduction()
		if err != nil {
			return err
		}
		defer l.Sync()

		// default is the get generator
		p := gen.NewGit(iface.WithLogger(l))

		// Fallback to archive if this mode is enabled
		if cfg.InitConfig.ArchiveMode {
			p = gen.NewArchive(cfg.InitConfig.URL, cfg.InitConfig.Folder)
		}

		// create root context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// clone the repository
		err = p.CloneWithContext(ctx, cfg.InitConfig.URL, cfg.InitConfig.Folder)
		if err != nil {
			return err
		}

		// run plugins ...
		exec := plugin.NewExecutor()
		err = exec.ExecWithContext(ctx)
		if err != nil {
			return err
		}

		return nil
	},
}
