package config

import (
	"os"

	"github.com/andersnormal/picasso/templates"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	// CfgFile
	File string `envconfig:"PICASSO_CONFIG_FILE" default:".picasso.yaml"`

	// LogFormat
	LogFormat string `envconfig:"PICASSO_CONFIG_LOG_FORMAT" default:""`

	// LogLevel
	LogLevel log.Level

	// Verbose
	Verbose bool `envconfig:"PICASSO_VERBOSE"`

	// FileMode
	FileMode os.FileMode
}

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

// Tasks ...
type Tasks map[string]*Task

// Watch ...
type Watch map[string]interface{}

// Vars ...
type Vars map[string][]Var

// Var
type Var string

// Settings ...
type Settings struct {
	Version string
	Author  string
	Project string
	Tasks   Tasks
	Vars    Vars
}
