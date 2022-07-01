package spec

import (
	"reflect"
	"strings"

	"github.com/andersnormal/picasso/pkg/plugin"
	"github.com/go-playground/validator/v10"
)

// DefaultFilename ...
const (
	DefaultFilename = ".picasso.yml"
)

// Spec ...
type Spec struct {
	// Spec ...
	Spec string `validate:"required" yaml:"spec"`
	// Version ...
	Version string `validate:"required" yaml:"version"`
	// Description ...
	Description string `yaml:"description"`
	// Authors ...
	Authors Authors `validate:"required" yaml:"authors"`
	// Homepage ...
	Homepage string `yaml:"homepage"`
	// License ...
	License string `yaml:"license"`
	// Repository ...
	Repository string `yaml:"repository"`
	// Plugins ...
	Plugins Plugins `yaml:"plugins"`
	// Tasks ...
	Tasks Tasks `yaml:"tasks"`
	// Generators ...
	Generators Generators `yaml:"generators"`
	// Template ...
	Template Template `yaml:"template"`
}

// Validate ..
func (s *Spec) Validate() error {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(s)
	if err != nil {
		return err
	}

	return v.Struct(s)
}

// Proto ...
func (s *Spec) Proto() *plugin.Spec {
	spec := &plugin.Spec{}

	return spec
}

// Authors ...
type Authors []string

// Plugins ...
type Plugins []string

// Tasks ...
type Tasks map[string]Task

// Generators ...
type Generators map[string]Generator

// Task ...
type Task struct {
	Id          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Commands    Commands  `yaml:"commands"`
	Disabled    bool      `yaml:"disabled"`
	Default     bool      `yaml:"default"`
	DependsOn   DependsOn `yaml:"depends_on"`
	Vars        Vars      `yaml:"vars"`
	Watch       Watch     `yaml:"watch"`
}

// Template ...
type Template struct {
	// Includes ...
	Includes Includes `yaml:"includes"`
	// Excludes ...
	Excludes Excludes `yaml:"excludes"`
	// Inputs ...
	Inputs Inputs `yaml:"inputs"`
	// Name ...
	Name string `yaml:"name"`
	// Description ...
	Description string `yaml:"description"`
	// File ...
	File string `yaml:"file"`
	// Out ...
	Out string `yaml:"out"`
	// Vars ...
	Vars Vars `yaml:"vars"`
}

// Watch ...
type Watch struct {
	Paths   Paths   `yaml:"paths,omitempty"`
	Ignores Ignores `yaml:"ignores,omitempty"`
}

// Paths ...
type Paths []string

// Ignore ...
type Ignores []string

// Vars ...
type Vars map[string]string

// DependsOn ...
type DependsOn []string

// Commands ...
type Commands []Command

// Command ...
type Command string

// Generator ...
type Generator struct {
	Id          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Path        string   `yaml:"path"`
	Includes    Includes `yaml:"includes"`
	Excludes    Excludes `yaml:"excludes"`
	Inputs      Inputs   `yaml:"inputs"`
}

// Inputs ...
type Inputs []Input

// Input ...
type Input struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Prompt string `yaml:"prompt"`
	Regex  string `yaml:"regex"`
}

// Includes ...
type Includes []string

// Excludes ...
type Excludes []string

// Plugin ...
type Plugin struct {
	Id          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
}
