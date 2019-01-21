package cmd

import (
	"fmt"
	"github.com/andersnormal/picasso/templates"
	"path"

	"github.com/andersnormal/picasso/gen"
	"github.com/andersnormal/picasso/settings"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	readmePackr = packr.New("readme", "../templates/readme")
)

func init() {
	Create.Flags().StringVar(&defaults.Author, "author", defaults.Author, "author")
	Create.Flags().StringVar(&defaults.Project, "project", defaults.Project, "project")
}

var Create = &cobra.Command{
	Use:   "create",
	Short: "creates a new project",
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
			},
		}

		// use README generator
		if err := generate(readmePackr, defaults, gopts...); err != nil {
			return err
		}

		// settings opts
		sopts := []settings.Opt{func(o *settings.Opts) {
			o.File = path.Join(cwd, settings.DefaultFile)
			o.FileMode = cfg.FileMode
		}}

		// read in config
		s := settings.New(sopts...)
		if err = s.Write(defaults); err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

func generate(b *packr.Box, gc gen.Context, opts ...gen.Opt) error {
	g := gen.NewGenerator(b, gc, opts...)

	if err := g.Write(); err != nil {
		return err
	}

	return nil
}
