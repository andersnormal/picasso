package executr

import (
	"context"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"
	"github.com/andersnormal/picasso/pkg/templr"
	"github.com/fsnotify/fsnotify"

	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// New ...
func New(opts ...Opt) Executr {
	options := &Opts{
		Env: make(Env),
	}
	options.Configure(opts...)

	e := new(exectur)
	e.opts = options

	return e
}

// Stdin ...
func (e *exectur) Stdin() io.Reader {
	return e.opts.Stdin
}

// Stdout ...
func (e *exectur) Stdout() io.Writer {
	return e.opts.Stdout
}

// Stderr ...
func (e *exectur) Stderr() io.Writer {
	return e.opts.Stderr
}

// Run ...
func (e *exectur) Run(ctx context.Context, exec Exec) error {
	timeout := time.Duration(time.Nanosecond * math.MaxInt)

	if !exec.Watch {
		timeout = time.Duration(e.opts.Timeout)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := e.runCmd(ctx, exec.WorkingDir, exec.Task.Commands)
	if err != nil {
		return err
	}

	err = e.genTemplates(exec.Task.Templates)
	if err != nil {
		return err
	}

	if !exec.Watch {
		return nil
	}

	fs, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer fs.Close()

	for _, p := range exec.Task.Watch.Paths {
		if err := fs.Add(p); err != nil {
			return err
		}
	}

Loop:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-fs.Events:
			for _, p := range exec.Task.Watch.Ignores {
				if strings.Contains(event.Name, p) {
					continue Loop
				}
			}

			err := e.runCmd(ctx, exec.WorkingDir, exec.Task.Commands)
			if err != nil {
				return err
			}
		case err := <-fs.Errors:
			return err
		}
	}
}

func (e *exectur) genTemplates(tt spec.Templates) error {
	fields, err := templr.DefaultFields()
	if err != nil {
		return err
	}

	t := templr.New(templr.WithFields(fields))

	for _, tpl := range tt {
		err := t.ParseFile(tpl.File, tpl.Out)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *exectur) runCmd(ctx context.Context, dir spec.WorkingDir, cmds []spec.Command) error {
	fields, err := templr.DefaultFields()
	if err != nil {
		return err
	}

	t := templr.New(templr.WithFields(fields))

	for _, cmd := range cmds {
		s := t.Parse(string(cmd))

		p, err := syntax.NewParser().Parse(strings.NewReader(s), "")
		if err != nil {
			return err
		}

		r, err := interp.New(
			interp.Dir(dir.String()),
			interp.Env(expand.ListEnviron(append(os.Environ(), e.opts.Env.Strings()...)...)),

			interp.Module(interp.DefaultExec),
			interp.Module(interp.OpenDevImpls(interp.DefaultOpen)),

			interp.StdIO(e.opts.Stdin, e.opts.Stdout, e.opts.Stderr),
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
