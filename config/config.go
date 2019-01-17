package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var C = NewConfig()

func NewConfig() *Config {
	c := &Config{
		LogLevel: log.WarnLevel,
	}
	envconfig.Process("", c)
	return c
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	// enable verbose output
	cmd.PersistentFlags().BoolVar(&c.Verbose, "verbose", c.Verbose, "enable verbose output")
}
