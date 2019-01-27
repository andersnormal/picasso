package task

import (
	"github.com/andersnormal/picasso/templates"
)

// Tasks ...
type Tasks map[string]*Task

// Task ...
type Task struct {
	Cmds      []Cmd
	Paths     []string
	Desc      string
	Vars      Vars
	Templates templates.Templates
}

// Cmd ...
type Cmd string

// Vars ...
type Vars map[string][]Var

// Var
type Var string
