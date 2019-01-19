package config

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	// CfgFile
	CfgFile string `envconfig:"PICASSO_CONFIG_FILE" default:""`

	// LogFormat
	LogFormat string `envconfig:"PICASSO_CONFIG_LOG_FORMAT" default:""`

	// LogLevel
	LogLevel log.Level

	// Verbose
	Verbose bool `envconfig:"PICASSO_VERBOSE"`
}
