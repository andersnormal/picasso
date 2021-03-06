package cmd

import (
	"context"

	"github.com/andersnormal/picasso/pkg/providers"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialized a new project from archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := providers.NewArchive(cfg.InitConfig.URL, cfg.InitConfig.Folder)

		// create root context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// noop
		return p.CloneWithContext(ctx)
	},
}
