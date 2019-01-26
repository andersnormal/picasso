package cmd

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/andersnormal/picasso/config"
	"github.com/andersnormal/picasso/executr"
	s "github.com/andersnormal/picasso/settings"
	"github.com/andersnormal/picasso/templates"
	"github.com/andersnormal/picasso/templr"
	"github.com/andersnormal/picasso/utils"

	"github.com/spf13/cobra"
)

func init() {}

var (
	hooks = []string{"preRun", "postRun"}
)

func init() {
	Build.Flags().StringVar(&settings.Author, "author", settings.Author, "author")
	Build.Flags().StringVar(&settings.Project, "project", settings.Project, "project")
}

var Build = &cobra.Command{
	Use:   "build",
	Short: "build a new project",
	RunE:  buildRunE,
}

func buildRunE(cmd *cobra.Command, args []string) error {
	var (
		hooks = []string{"preRun", "postRun"}
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// configure path
	cwd, err := cfg.Cwd()
	if err != nil {
		return err
	}

	// settings opts
	sopts := []s.Opt{func(o *s.Opts) {
		o.File = path.Join(cwd, cfg.File)
		o.FileMode = cfg.FileMode
	}}

	// new settings
	settings := config.NewSettings()
	ss := s.New(sopts...)
	if err := ss.Read(&settings); err != nil {
		return err
	}

	for _, hook := range hooks {
		task, ok := settings.Tasks[hook]
		if !ok {
			continue
		}

		for _, tmpl := range task.Templates {
			if err := writeTemplate(ctx, cwd, tmpl); err != nil {
				return err
			}
		}

		fmt.Printf("executing %s", task.Desc)

		err := execTask(ctx, task)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeTemplate(ctx context.Context, cwd string, tmpl *templates.Template) error {
	in := path.Join(cwd, tmpl.File)
	out := path.Join(cwd, tmpl.Output)

	t, err := utils.Stream(in)
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	topts := []templr.Opt{
		func(o *templr.Opts) {
			o.Vars = make(templr.Vars)
			for k, v := range tmpl.Vars {
				o.Vars[k] = templr.Var(v)
			}
		},
	}

	tr := templr.New(topts...)
	_, err = f.WriteString(tr.Parse(string(t)))
	if err != nil {
		return err
	}

	return nil
}

func execTask(ctx context.Context, task *config.Task) error {
	for _, cmd := range task.Cmds {
		opts := []executr.Opt{func(o *executr.Opts) {
			o.Cmd = string(cmd)
		}}

		e := executr.New(opts...)

		if err := e.Run(ctx); err != nil {
			return err
		}

	}

	return nil
}
