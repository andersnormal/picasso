package runner

import (
	"context"
	"fmt"

	"golang.org/x/exp/maps"
)

// Ctx ...
type Ctx struct {
	runner     *Runner
	cmd        Cmd
	workingDir WorkingDir
	vars       Values[string, string]
	env        Values[string, string]
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

// Values ...
type Values[K comparable, T any] map[K]Value[T]

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
