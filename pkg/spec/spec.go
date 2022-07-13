package spec

import (
	"fmt"
	"io/ioutil"
	"os"
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

// Environ ...
func (s *Spec) Environ() []string {
	if s.Env == nil {
		return nil
	}

	environ := os.Environ()

	for k, v := range s.Env {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}

	return environ
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
	all := make(map[string]bool)
	tt := make([]Task, 0)

	for _, name := range names {
		if _, exists := all[name]; exists {
			continue
		}

		t, ok := s.Tasks[name]
		if !ok {
			return nil, ErrTaskNotFound
		}

		for _, dep := range t.DependsOn {
			if _, exists := all[dep]; exists {
				continue
			}

			d, ok := s.Tasks[dep]
			if !ok {
				return nil, ErrTaskNotFound
			}
			tt = append(tt, d)
			all[d.Name] = true
		}

		all[name] = true
		tt = append(tt, t)
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
	If        string    `yaml:"if"`
	Default   bool      `yaml:"default"`
	DependsOn DependsOn `yaml:"depends-on"`
	Name      string    `yaml:"name"`
	Disabled  bool      `yaml:"disabled"`
	Env       Env       `yaml:"env"`
	Vars      Vars      `yaml:"vars"`
	Templates Templates `yaml:"template,omitempty"`

	Watch      Watch      `yaml:"watch"`
	WorkingDir WorkingDir `yaml:"working-dir"`
	Steps      Steps      `yaml:"steps"`
}

// Step ...
type Step struct {
	Id               string            `yaml:"id"`
	Cmd              string            `yaml:"cmd"`
	Env              Env               `yaml:"env"`
	Vars             Vars              `yaml:"vars"`
	If               string            `yaml:"if"`
	TimeoutInSeconds int64             `yaml:"timeout-in-seconds"`
	Uses             string            `yaml:"uses"`
	With             map[string]string `yaml:"with"`
	WorkingDir       WorkingDir        `yaml:"working-dir"`
	ContinueOnError  bool              `yaml:"continue-on-error"`
	Templates        Templates         `yaml:"template,omitempty"`
}

// Steps ...
type Steps []Step

// Task ...
func (t *Task) Environ() []string {
	if t.Env == nil {
		return nil
	}

	environ := os.Environ()

	for k, v := range t.Env {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}

	return environ
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
