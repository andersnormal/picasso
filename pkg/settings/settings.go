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

func New(opts ...Opt) Settings {
	options := &Opts{}

	var s = new(settings)
	s.opts = options

	_ = configure(s, opts...)

	return s
}

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

func configure(s *settings, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	if s.opts.File == "" {
		s.opts.File = DefaultFile
	}

	if s.opts.FileMode == 0 {
		s.opts.FileMode = DefaultFileMode
	}

	return nil
}
