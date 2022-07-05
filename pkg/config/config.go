package config

import (
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"
)

// Flags ...
type Flags struct {
	Version   bool
	Generator string
	List      bool
	Validate  bool
	Silent    bool
	Help      bool
	Init      bool
	Force     bool
	Verbose   bool
	Dry       bool
	Env       []string
	Timeout   time.Duration
}

// Config ...
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
	// Flags ...
	Flags Flags
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File
}

// InitConfig ...
type InitConfig struct {
	// Folder ...
	Folder string
	// URL ...
	URL string
	// ArchiveMode ...
	ArchiveMode bool
}

// New ...
func New() *Config {
	return &Config{
		File:         ".picasso.yml",
		KillSignal:   syscall.SIGINT,
		LogFormat:    "text",
		LogLevel:     "warn",
		ReloadSignal: syscall.SIGHUP,
		TermSignal:   syscall.SIGTERM,
		Stdin:        os.Stdin,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
	}
}

// InitDefaultConfig() ...
func (c *Config) InitDefaultConfig() error {
	cwd, err := c.Cwd()
	if err != nil {
		return err
	}
	c.File = filepath.Join(cwd, c.File)

	return nil
}

// Cwd ...
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

// SpecFile ...
func (c *Config) LoadSpec() (*spec.Spec, error) {
	s, err := spec.Load(c.File)
	if err != nil {
		return nil, err
	}

	return s, nil
}
