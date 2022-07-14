package runner

import (
	"context"
	"io"
	"math"
	"strings"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"
	"github.com/andersnormal/picasso/pkg/templr"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// AddStringSlice ...
func AddStringSlice(vars []string) RunFunc {
	return func(c *Ctx) error {
		for _, v := range vars {
			kv := strings.Split(v, "=")
			if len(kv) != 2 {
				c.Vars().Add(kv[0], "")
				continue
			}
			c.Vars().Add(kv[0], kv[1])
		}

		return nil
	}
}

// AddVars...
func AddVars(vars spec.Vars) RunFunc {
	return func(c *Ctx) error {
		for k, v := range vars {
			c.Vars().Add(k, v)
		}

		return nil
	}
}

// RunTask ...
func RunTask(t spec.Task, cwd string) RunFunc {
	return func(c *Ctx) error {
		if t.Disabled {
			return nil
		}

		for _, s := range t.Steps {
			cmds := strings.Split(s.Cmd, "\n")
			timeout := time.Duration(time.Nanosecond * math.MaxInt)

			if s.TimeoutInSeconds > 0 {
				timeout = time.Duration(s.TimeoutInSeconds) * time.Second
			}

			if err := runCmd(c.Context(), timeout, c.WorkingDir(), c.runner.Stdin(), c.runner.Stdout(), c.runner.Stderr(), t.Environ(), cmds...); err != nil {
				if s.ContinueOnError {
					continue
				}

				return err
			}
		}

		return nil
	}
}

func runCmd(ctx context.Context, timeout time.Duration, dir WorkingDir, stdin io.Reader, stdout io.Writer, stderr io.Writer, env []string, cmds ...string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	fields, err := templr.DefaultFields()
	if err != nil {
		return err
	}

	_ = templr.New(templr.WithFields(fields))

	for _, cmd := range cmds {
		p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
		if err != nil {
			return err
		}

		r, err := interp.New(
			interp.Dir(string(dir)),
			interp.Env(expand.ListEnviron(env...)),

			interp.Module(interp.DefaultExec),
			interp.Module(interp.OpenDevImpls(interp.DefaultOpen)),

			interp.StdIO(stdin, stdout, stderr),
		)
		if err != nil {
			return err
		}

		err = r.Run(ctx, p)
		if err != nil {
			return err
		}
	}

	return nil
}