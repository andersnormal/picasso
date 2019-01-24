package config

import (
// "github.com/andersnormal/picasso/templr"
)

func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var settings struct {
		Version string
		Author  string
		Project string
		Build   Build
		Vars    Vars
	}
	if err := unmarshal(&settings); err != nil {
		return err
	}

	s.Version = settings.Version
	s.Build = settings.Build
	s.Vars = settings.Vars
	s.Build = settings.Build

	return nil
}
