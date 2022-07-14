package runner

import (
	"context"
	"fmt"
	"strings"

	"github.com/andersnormal/picasso/pkg/spec"
	"golang.org/x/exp/maps"
)

// Ctx ...
type Ctx struct {
	funcs      []RunFunc
	idx        int
	cmd        Cmd
	env        Env
	runner     *Runner
	vars       Vars
	workingDir WorkingDir
	task       spec.Task
}

// WorkingDir ..
type WorkingDir string

// WorkingDir ..
func (c *Ctx) WorkingDir() WorkingDir {
	return c.workingDir
}

// Cmd ...
type Cmd string

// Cmd ...
func (c *Ctx) Cmd() Cmd {
	return c.cmd
}

// Runner ...
func (c *Ctx) Runner() *Runner {
	return c.runner
}

// Task ...
func (c *Ctx) Task() spec.Task {
	return c.task
}

// Reset ...
func (c *Ctx) Reset() {
	c.env = make(Env)
	c.vars = make(Vars)
	c.cmd = ""
	c.workingDir = ""
}

// Next ...
func (c *Ctx) Next() error {
	c.idx++
	if c.idx < len(c.funcs) {
		if err := c.funcs[c.idx](c); err != nil {
			return err
		}
	}

	return nil
}

// Context ...
func (c *Ctx) Context() context.Context {
	return c.runner.Context()
}

// Vars ...
func (c *Ctx) Vars() Values[string, string] {
	return c.vars.Clone()
}

// Env ...
func (c *Ctx) Env() []string {
	env := make([]string, len(c.env))
	for k, v := range c.env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env
}

// Vars ...
type Vars = Values[string, string]

// Env ...
type Env = Values[string, string]

// Values ...
type Values[K comparable, T any] map[K]Value[T]

// NewFromSlice ...
func NewFromSlice(s []string) Values[string, string] {
	vv := make(Values[string, string])
	for _, v := range s {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			vv.Add(kv[0], "")
			continue
		}
		vv.Add(kv[0], kv[1])
	}

	return vv
}

// NewFromMap ...
func NewFromMap(m map[string]string) Values[string, string] {
	vv := make(Values[string, string])
	for k, v := range m {
		vv.Add(k, v)
	}

	return vv
}

// Add ..
func (vv Values[K, T]) Add(key K, value T) {
	vv[key] = Value[T]{val: value}
}

// Copy ...
func (vv Values[K, T]) Copy(m Values[K, T]) {
	maps.Copy(vv, m)
}

// Clone ...
func (vv Values[K, T]) Clone() Values[K, T] {
	return maps.Clone(vv)
}

// Clear ...
func (vv Values[K, T]) Clear() {
	maps.Clear(vv)
}

// Value ...
type Value[T any] struct {
	val T
}

// Value ...
func (v *Value[T]) Value() T {
	return v.val
}
