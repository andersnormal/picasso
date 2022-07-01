package cmd

import (
	"github.com/andersnormal/picasso/pkg/settings"
	"github.com/andersnormal/picasso/pkg/spec"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a task file",
	RunE: func(cmd *cobra.Command, args []string) error {
		spath, err := cfg.Settings()
		if err != nil {
			return err
		}

		spec := spec.Spec{}
		s := settings.New(settings.WithFile(spath))
		err = s.Read(&spec)
		if err != nil {
			return err
		}

		return spec.Validate()
	},
}
