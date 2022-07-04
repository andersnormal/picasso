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
	Version bool
	Silent  bool
	Help    bool
	Init    bool
	Force   bool
	Verbose bool
	Dry     bool
	Env     []string
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
	// InitConfig ...
	InitConfig InitConfig
	// RunConfig ...
	RunConfig RunConfig
	// Plugins ...
	Plugins []string

	// Flags ...
	Flags Flags
	// Stdin ...
	Stdin *os.FileInfo
	// Stdout ...
	Stdout *os.FileInfo
	// Stderr ...
	Stderr *os.FileInfo
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

// RunConfig ...
type RunConfig struct {
	// Env ...
	Env []string
	// TImeout ...
	Timeout time.Duration
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
	Vars    Vars
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
		Verbose:      false,
		InitConfig:   InitConfig{Folder: "", URL: ""},
		RunConfig:    RunConfig{Env: []string{}, Timeout: time.Second * 300},
	}
}

// InitDefaultConfig() ...
func (c *Config) InitDefaultConfig() error {
	cwd, err := c.Cwd()
	if err != nil {
		return err
	}
	c.File = filepath.Join(cwd, c.File)
	c.InitConfig.Folder = cwd

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
