package specs

import (
	"github.com/andersnormal/picasso/pkg/templates"
)

// Spec ...
type Spec struct {
	// Spec ...
	Spec string `validate:"required"`
	// Version ...
	Version string `validate:"required"`
	// Description ...
	Description string
	// Author ...
	Author string `validate:"required"`
	// Includes ...
	Includes []string `validate:"required_with=Excludes"`
	// include ...
	Excludes []string
	// Tasks ...
	Tasks Tasks
	// Plugins ...
	Plugins Plugins
}

// Plugins ...
type Plugins []string

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
	Deps      Deps                `yaml:"deps,omitempty"`
	Desc      string              `yaml:"desc,omitempty"`
	Disable   bool                `yaml:"disable,omitempty"`
	Default   bool                `yaml:"default,omitempty"`
	Paths     Paths               `yaml:"paths,omitempty"`
	Templates templates.Templates `yaml:"templates,omitempty"`
	Vars      Vars                `yaml:"vars,omitempty"`
	Watch     Watch               `yaml:"watch,omitempty"`

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
