package cmd

import (
	"context"

	git "github.com/andersnormal/picasso/pkg/init"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                "init",
	Short:              "Initialized a new project",
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		// default is the git generator
		p := git.NewGit()

		// Fallback to archive if this mode is enabled
		if cfg.InitConfig.ArchiveMode {
			p = git.NewArchive(cfg.InitConfig.URL, cfg.InitConfig.Folder)
		}

		// create root context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// clone the repository
		err := p.CloneWithContext(ctx, cfg.InitConfig.URL, cfg.InitConfig.Folder)
		if err != nil {
			return err
		}

		return nil
	},
}
