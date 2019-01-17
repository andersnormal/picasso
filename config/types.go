package config

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	// LogFormat
	LogFormat string `envconfig:"PICASSO_LOG_FORMAT" default:""`

	// LogLevel
	LogLevel log.Level

	// Verbose
	Verbose bool `envconfig:"PICASSO_VERBOSE"`
}
