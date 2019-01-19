package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	ConfigFile = ".picasso"
)

func New() *Config {
	c := &Config{
		LogLevel: log.WarnLevel,
	}
	envconfig.Process("", c)
	return c
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	// enable verbose output
	cmd.PersistentFlags().BoolVar(&c.Verbose, "verbose", c.Verbose, "enable verbose output")

	// enable verbose output
	cmd.PersistentFlags().StringVar(&c.CfgFile, "config", c.CfgFile, "config file (default is .picasso.yaml)")
}
