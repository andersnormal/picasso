package cmd

import (
	"reflect"
	"strings"

	"github.com/andersnormal/picasso/pkg/settings"
	"github.com/andersnormal/picasso/pkg/specs"

	"github.com/go-playground/validator/v10"
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

		spec := specs.Spec{}
		s := settings.New(settings.WithFile(spath))
		err = s.Read(&spec)
		if err != nil {
			return err
		}

		v := validator.New()
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		err = v.Struct(spec)
		if err != nil {
			return err
		}

		return nil
	},
}
