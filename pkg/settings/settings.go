package settings

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andersnormal/picasso/pkg/utils"

	"gopkg.in/yaml.v2"
)

const (
	DefaultFile     = ".picasso.yaml"
	DefaultFileMode = 0755
)

type settings struct {
	opts *Opts
}

type Settings interface {
	Read(in interface{}) error
	Write(out interface{}) error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	File     string
	FileMode os.FileMode
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}

	if o.File == "" {
		o.File = DefaultFile
	}

	if o.FileMode == 0 {
		o.FileMode = DefaultFileMode
	}
}

// New ...
func New(opts ...Opt) Settings {
	options := &Opts{}
	options.Configure(opts...)

	var s = new(settings)
	s.opts = options

	return s
}

// Read ...
func (s *settings) Read(in interface{}) error {
	if _, err := os.Stat(s.opts.File); os.IsNotExist(err) {
		return fmt.Errorf("settings error: %w", err)
	}

	c, err := utils.Stream(s.opts.File)
	if err != nil {
		return fmt.Errorf("settings error: %w", err)
	}

	err = yaml.Unmarshal(c, in)
	if err != nil {
		return fmt.Errorf("settings error: %w", err)
	}

	return nil
}

// Write ...
func (s *settings) Write(out interface{}) error {
	y, err := yaml.Marshal(out)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.opts.File, y, s.opts.FileMode)
	if err != nil {
		return err
	}

	return nil
}

// WithFile ...
func WithFile(f string) Opt {
	return func(o *Opts) {
		o.File = f
	}
}
