package cmd

import (
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a task file",
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := cfg.LoadSpec()
		if err != nil {
			return err
		}

		return spec.Validate()
	},
}
