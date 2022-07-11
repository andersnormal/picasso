package spec

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/andersnormal/picasso/pkg/templr"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

var (
	ErrTaskNotFound = fmt.Errorf("task not found")
)

// DefaultFilename ...
const (
	DefaultFilename = ".picasso.yml"
)

// Spec ...
type Spec struct {
	// Spec ...
	Spec int `validate:"required" yaml:"spec"`
	// Version ...
	Version string `validate:"required" yaml:"version,omitempty"`
	// Description ...
	Description string `yaml:"description,omitempty"`
	// Authors ...
	Authors Authors `validate:"required" yaml:"authors,omitempty"`
	// Homepage ...
	Homepage string `yaml:"homepage,omitempty"`
	// License ...
	License string `yaml:"license,omitempty"`
	// Repository ...
	Repository string `yaml:"repository,omitempty"`
	// Plugins ...
	Plugins Plugins `yaml:"plugins,omitempty"`
	// Tasks ...
	Tasks Tasks `yaml:"tasks"`
	// Vars ...
	Vars Vars `yaml:"vars"`
	// Env ...
	Env Env `yaml:"env"`
}

// Fields ...
func (s *Spec) Fields() templr.Fields {
	fields := templr.Fields{
		"Spec":        s.Spec,
		"Version":     s.Version,
		"Description": s.Description,
		"Authors":     s.Authors,
		"License":     s.License,
		"Repository":  s.Repository,
	}

	return fields
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

// Load ...
func Load(file string) (*Spec, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(f, &spec)
	if err != nil {
		return nil, err
	}

	return &spec, nil
}

// Default ...
func (s *Spec) Default() []Task {
	tt := make([]Task, 0)

	for _, t := range s.Tasks {
		if t.Default {
			tt = append(tt, t)
		}
	}

	return tt
}

// Find ...
func (s *Spec) Find(names []string) ([]Task, error) {
	tt := make([]Task, 0, len(names))

	for _, name := range names {
		found := false
		for k, t := range s.Tasks {
			if k == name {
				tt = append(tt, t)
				found = true
			}
		}

		if !found {
			return nil, ErrTaskNotFound
		}
	}

	return tt, nil
}

// Authors ...
type Authors []string

// Plugins ...
type Plugins []Plugin

// Tasks ...
type Tasks map[string]Task

// Task ...
type Task struct {
	Id string `yaml:"id"`

	Commands    Commands   `yaml:"cmd"`
	Default     bool       `yaml:"default"`
	DependsOn   DependsOn  `yaml:"depends_on"`
	Description string     `yaml:"description"`
	Disabled    bool       `yaml:"disabled"`
	Env         Env        `yaml:"env"`
	Name        string     `yaml:"name"`
	Templates   Templates  `yaml:"template,omitempty"`
	Vars        Vars       `yaml:"vars"`
	Watch       Watch      `yaml:"watch"`
	WorkingDir  WorkingDir `yaml:"working_dir"`
}

// WorkingDir ...
type WorkingDir string

// String ...
func (w WorkingDir) String() string {
	return string(w)
}

// Templates ...
type Templates []Template

// Template ...
type Template struct {
	File string `yaml:"file"`
	Out  string `yaml:"out"`
	Vars Vars   `yaml:"var"`
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

// Merge ...
func (v Vars) Merge(vars Vars) {
	maps.Copy(v, vars)
}

// Env ...
type Env map[string]string

// DependsOn ...
type DependsOn []string

// Commands ...
type Commands []Command

// Command ...
type Command string

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
