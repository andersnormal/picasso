package context

import (
	"os"
	"text/template"

	"github.com/andersnormal/picasso/gen"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _ gen.Context = (*createContext)(nil)

func NewCreateContext() gen.Context {
	return &createContext{}
}

type createContext struct {
	Project string
	Author  string
}

func (c *createContext) Flags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.Project, "project", c.Project, "project name")
	cmd.Flags().StringVar(&c.Author, "author", c.Author, "author")

	viper.BindPFlag("project", cmd.Flags().Lookup("project"))
	viper.BindPFlag("author", cmd.Flags().Lookup("author"))
}

func (c *createContext) Execute(t *template.Template, f *os.File) error {
	return t.Execute(f, c)
}
