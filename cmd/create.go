package cmd

import (
	"fmt"

	"github.com/andersnormal/picasso/context"
	"github.com/andersnormal/picasso/gen"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var (
	gc gen.Context
)

func init() {
	gc := context.NewCreateContext()
	gc.Flags(Create)
}

var Create = &cobra.Command{
	Use:   "create",
	Short: "creates a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("creating project")

		// use README generator
		if err := generate(packr.New("readme", "../templates/readme"), gc); err != nil {
			return err
		}

		return nil
	},
}

func generate(b *packr.Box, gc gen.Context) error {
	g := gen.NewGenerator(b, gc)

	if err := g.Write(); err != nil {
		return err
	}

	return nil
}
