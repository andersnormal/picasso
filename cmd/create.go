package cmd

import (
	"fmt"
	"path"

	"github.com/andersnormal/picasso/config"
	"github.com/andersnormal/picasso/gen"
	s "github.com/andersnormal/picasso/settings"
	"github.com/andersnormal/picasso/templates"
	"github.com/andersnormal/picasso/templr"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var (
	readmePackr = packr.New("readme", "../templates/readme")
	basicPackr  = packr.New("readme", "../templates/basic")
)

func generateCreate(cwd string) *cobra.Command {
	// new settings
	settings := config.NewSettings()

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("creating project")

			// configure path
			cwd, err := cfg.Cwd()
			if err != nil {
				return err
			}

			if len(args) > 0 && path.IsAbs(args[0]) {
				cwd = args[0]
			} else if len(args) > 0 && !path.IsAbs(args[0]) {
				cwd = path.Join(cwd, args[0])
			}

			// opts for generator
			gopts := []gen.Opt{
				func(o *gen.Opts) {
					o.Dir = cwd
					o.Templates = templates.Packr
					o.Vars = templr.Vars{
						"author":  templr.Var(settings.Author),
						"project": templr.Var(settings.Project),
					}
				},
			}

			// use README generator
			if err := generate(readmePackr, settings, gopts...); err != nil {
				return err
			}

			// settings opts
			sopts := []s.Opt{func(o *s.Opts) {
				o.File = path.Join(cwd, s.DefaultFile)
				o.FileMode = cfg.FileMode
			}}

			s := s.New(sopts...)

			if err = s.Write(settings); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&settings.Author, "author", settings.Author, "author")
	cmd.Flags().StringVar(&settings.Project, "project", settings.Project, "project")

	return cmd
}

func generate(b *packr.Box, gc gen.Context, opts ...gen.Opt) error {
	g := gen.NewGenerator(b, gc, opts...)

	if err := g.Write(); err != nil {
		return err
	}

	return nil
}
