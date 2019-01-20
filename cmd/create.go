package cmd

import (
	"fmt"
	"path"

	"github.com/andersnormal/picasso/gen"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Create.Flags().String("name", "", "name")
	Create.Flags().String("author", "", "author")

	viper.BindPFlag("name", Create.Flags().Lookup("name"))
	viper.BindPFlag("author", Create.Flags().Lookup("author"))
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

		// try to create path
		if err := gen.MkdirAll(cwd, cfg.FileMode); err != nil {
			return err
		}

		// opts for generator
		opts := []gen.Opt{func(o *gen.Opts) {
			o.Dir = cwd
		}}

		// construct context
		gc := struct {
			Name   string
			Author string
		}{
			Name:   viper.GetString("name"),
			Author: viper.GetString("author"),
		}

		// use README generator
		if err := generate(packr.New("readme", "../templates/readme"), gc, opts...); err != nil {
			return err
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
