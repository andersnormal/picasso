package cmd

import (
	"context"

	"github.com/andersnormal/picasso/pkg/providers"
	"github.com/andersnormal/picasso/pkg/providers/iface"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialized a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := zap.NewProduction()
		if err != nil {
			return err
		}

		p := providers.NewGit(iface.WithLogger(l))

		// Fallback to archive if this mode is enabled
		if cfg.InitConfig.ArchiveMode {
			p = providers.NewArchive(cfg.InitConfig.URL, cfg.InitConfig.Folder)
		}

		// create root context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// noop
		return p.CloneWithContext(ctx, cfg.InitConfig.URL, cfg.InitConfig.Folder)
	},
}
