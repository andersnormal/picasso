package executr

import (
	"context"
	"fmt"
	"io"
	"time"
)

type Executr interface {
	Run(ctx context.Context) error
	Env() []string
}

type Env map[string]string

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

type Opt func(*Opts)

type Opts struct {
	Cmd     string
	Dir     string
	Env     Env
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Timeout time.Duration
}
