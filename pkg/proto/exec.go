package proto

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"google.golang.org/protobuf/proto"
)

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
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

type executor struct {
	opts *Opts
}

// Executor ...
type Executor interface {
	// ExecWithContext ...
	ExecWithContext(context.Context, string, *PluginRequest) error
}

// NewExecutor ...
func NewExecutor(opts ...Opt) Executor {
	options := &Opts{}
	options.Configure(opts...)

	e := new(executor)
	e.opts = options

	return e
}

// ExecWithContext ...
func (e *executor) ExecWithContext(ctx context.Context, p string, req *PluginRequest) error {
	exec := exec.CommandContext(ctx, p)
	stdin, err := exec.StdinPipe()
	if err != nil {
		return err
	}

	m, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		_, err = io.Copy(stdin, bytes.NewReader(m))
		if err != nil {
			fmt.Println(err)
		}
	}()

	bb, err := exec.CombinedOutput()
	if err != nil {
		return err
	}

	msg := &PluginResponse{}
	err = proto.Unmarshal(bb, msg)
	if err != nil {
		return err
	}

	if msg.Error != "" {
		return fmt.Errorf(msg.Error)
	}

	return nil
}
