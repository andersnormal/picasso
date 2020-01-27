package cmd

import (
	"github.com/andersnormal/picasso/pkg/config"
	"github.com/spf13/cobra"
)

var run = &cobra.Command{
	Use:   "run [command]",
	Short: "Run task from your task file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.ErrNoDefaultTask
	},
}
