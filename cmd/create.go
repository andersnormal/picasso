package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/andersnormal/picasso/gen"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var _ gen.Context = (*context)(nil)
var gc *context

type context struct {
	Project string
}

func (c *context) Flags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.Project, "project", c.Project, "project name")
}

func (c *context) Execute(t *template.Template, f *os.File) error {
	return t.Execute(f, c)
}

func init() {
	gc = new(context)
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
