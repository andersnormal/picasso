package config

import (
	"os"
	"syscall"

	"github.com/andersnormal/picasso/pkg/task"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Config struct {
	// Verbose toggles the verbosity
	Verbose bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
	// ReloadSignal
	ReloadSignal syscall.Signal
	// TermSignal
	TermSignal syscall.Signal
	// KillSignal
	KillSignal syscall.Signal
	// CfgFile
	File string
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

func New() *Config {
	return &Config{
		File:         ".picasso.yaml",
		KillSignal:   syscall.SIGINT,
		LogFormat:    "text",
		LogLevel:     "warn",
		ReloadSignal: syscall.SIGHUP,
		TermSignal:   syscall.SIGTERM,
		Verbose:      false,
	}
}

func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

// SetupLogger prepares the logger instance
func (c *Config) SetupLogger() {
	switch c.LogFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.JSONFormatter{})
	}

	if c.Verbose {
		c.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := log.ParseLevel(c.LogLevel); err == nil {
		log.SetLevel(level)
	}
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	// enable verbose output
	cmd.PersistentFlags().BoolVar(&c.Verbose, "verbose", false, "enable verbose output")
	// set log format
	cmd.PersistentFlags().StringVar(&c.LogFormat, "log-format", "text", "log format")
	// set log level
	cmd.PersistentFlags().StringVar(&c.LogLevel, "log-level", "warn", "log level")
	// enable verbose output
	cmd.PersistentFlags().StringVar(&c.File, "config", c.File, "config file")
}
