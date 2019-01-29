package task

import (
	"github.com/andersnormal/picasso/templates"
)

// Tasks ...
type Tasks map[string]*Task

// Watch
type Watch struct {
	Paths  Paths    `yaml:"paths,omitempty"`
	Ignore []string `yaml:"ignore,omitempty"`
}

// Task ...
type Task struct {
	Cmds      []Cmd               `yaml:"cmds,omitempty"`
	Paths     Paths               `yaml:"paths,omitempty"`
	Deps      Deps                `yaml:"deps,omitempty"`
	Desc      string              `yaml:"desc,omitempty"`
	Vars      Vars                `yaml:"vars,omitempty"`
	Watch     Watch               `yaml:"watch,omitempty"`
	Templates templates.Templates `yaml:"templates,omitempty"`

	resolvedDeps []*Task
}

// Cmd ...
type Cmd string

// Paths
type Paths []string

// Deps
type Deps []string

// Vars ...
type Vars map[string][]Var

// Var
type Var string
