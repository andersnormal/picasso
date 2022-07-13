package executer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"
	"github.com/andersnormal/picasso/pkg/templr"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

var (
	ErrNoCmd   = errors.New("executer: no command to execute")
	ErrTimeout = errors.New("executer: command timeout")
)

// Executr ...
type Executr interface {
	// Run ...
	Run(ctx context.Context, exec Exec) error
}

// Exec ...
type Exec struct {
	WorkingDir spec.WorkingDir
	Watch      bool
	Env        spec.Env
	Task       spec.Task
}

// Cmd ...
type Cmd string

// Env ...
type Env map[string]string

// Strings ...
func (ev Env) Strings() []string {
	var env []string
	for k, v := range ev {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env
}

type exectur struct {
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Timeout time.Duration
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}

	if o.Stdin == nil {
		o.Stdin = os.Stdin
	}

	if o.Stdout == nil {
		o.Stdout = os.Stdout
	}

	if o.Stderr == nil {
		o.Stderr = os.Stderr
	}
}

// WithStdin ...
func WithStdin(r io.Reader) Opt {
	return func(o *Opts) {
		o.Stdin = r
	}
}

// WithStdout ...
func WithStdout(w io.Writer) Opt {
	return func(o *Opts) {
		o.Stdout = w
	}
}

// WithStderr ...
func WithStderr(w io.Writer) Opt {
	return func(o *Opts) {
		o.Stderr = w
	}
}

// WithTimeout ...
func WithTimeout(timeout time.Duration) Opt {
	return func(o *Opts) {
		o.Timeout = timeout
	}
}

// New ...
func New(opts ...Opt) Executr {
	options := &Opts{}
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
	for _, s := range exec.Task.Steps {
		cmds := strings.Split(s.Cmd, "\n")
		timeout := time.Duration(time.Nanosecond * math.MaxInt)

		if s.TimeoutInSeconds > 0 {
			timeout = time.Duration(s.TimeoutInSeconds) * time.Second
		}

		if err := e.runCmd(ctx, timeout, exec.WorkingDir, exec.Task.Environ(), cmds...); err != nil {
			if s.ContinueOnError {
				continue
			}

			return err
		}
	}

	// 	if !exec.Watch {
	// 		timeout = time.Duration(e.opts.Timeout)
	// 	}

	// 	ctx, cancel := context.WithTimeout(ctx, timeout)
	// 	defer cancel()

	// 	err := e.runCmd(ctx, exec.WorkingDir, exec.Task.Environ(), exec.Task.Commands)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = e.genTemplates(exec.Task.Templates)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if !exec.Watch {
	// 		return nil
	// 	}

	// 	fs, err := fsnotify.NewWatcher()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer fs.Close()

	// 	for _, p := range exec.Task.Watch.Paths {
	// 		if err := fs.Add(p); err != nil {
	// 			return err
	// 		}
	// 	}

	// Loop:
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return ctx.Err()
	// 		case event := <-fs.Events:
	// 			for _, p := range exec.Task.Watch.Ignores {
	// 				if strings.Contains(event.Name, p) {
	// 					continue Loop
	// 				}
	// 			}

	// 			err := e.runCmd(ctx, exec.WorkingDir, exec.Task.Environ(), exec.Task.Commands)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		case err := <-fs.Errors:
	// 			return err
	// 		}
	// 	}
	// }

	// func (e *exectur) genTemplates(tt spec.Templates) error {
	// 	fields, err := templr.DefaultFields()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	t := templr.New(templr.WithFields(fields))

	// 	for _, tpl := range tt {
	// 		err := t.ParseFile(tpl.File, tpl.Out)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	return nil
}

func (e *exectur) runCmd(ctx context.Context, timeout time.Duration, dir spec.WorkingDir, env []string, cmds ...string) error {
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
			interp.Dir(dir.String()),
			interp.Env(expand.ListEnviron(env...)),

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
