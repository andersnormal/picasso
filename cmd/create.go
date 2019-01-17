package cmd

import (
	"fmt"

	"github.com/andersnormal/picassso/utils"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

func init() {
}

var Create = &cobra.Command{
	Use:   "create",
	Short: "creates a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("creating project")

		// use readme generator
		if err := generate(packr.New("readme", "../templates/readme")); err != nil {
			return err
		}

		return nil
	},
}

func generate(b *packr.Box) error {
	g := utils.NewGenerator(b, &utils.GeneratorContext{})

	if err := g.Write(); err != nil {
		return err
	}

	return nil
}
