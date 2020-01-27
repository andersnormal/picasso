package cmd

import (
	"github.com/andersnormal/picasso/pkg/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [url]",
	Short: "Initialized a new project from archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.ErrNoDefaultTask
	},
}

func init() {
	initCmd.Flags().String("folder", "", "folder")
}
