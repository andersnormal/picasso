package config

import (
	"os"

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
