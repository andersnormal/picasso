package config

import (
	"os"

	"github.com/andersnormal/picasso/task"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	// CfgFile
	File string `envconfig:"PICASSO_CONFIG_FILE" default:".picasso.yaml"`

	// LogFormat
	LogFormat string `envconfig:"PICASSO_CONFIG_LOG_FORMAT" default:"text"`

	// LogLevel
	LogLevel log.Level

	// Verbose
	Verbose bool `envconfig:"PICASSO_VERBOSE"`

	// FileMode
	FileMode os.FileMode
}

// Vars ...
type Vars map[string][]Var

// Var
type Var string

// Settings ...
type Settings struct {
	Version string
	Author  string
	Project string
	Tasks   task.Tasks
	Vars    Vars
}
