package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func New() *Config {
	c := &Config{
		LogLevel: log.WarnLevel,
	}
	envconfig.Process("", c)
	return c
}

func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	// enable verbose output
	cmd.PersistentFlags().BoolVar(&c.Verbose, "verbose", c.Verbose, "enable verbose output")

	// enable verbose output
	cmd.PersistentFlags().StringVar(&c.File, "config", c.File, "config file (default is .picasso.yaml)")
}
