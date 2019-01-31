package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andersnormal/picasso/task"
)

func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) Task(n string) (*task.Task, error) {
	task, ok := s.Tasks[n]
	if !ok {
		return nil, ErrNoWatch
	}

	return task, nil
}

func (s *Settings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var (
		defaultTask bool

		settings struct {
			Version string
			Author  string
			Project string
			Tasks   task.Tasks
			Vars    Vars
		}
	)

	if err := unmarshal(&settings); err != nil {
		return err
	}

	// try to resolve tasks
	for name, task := range settings.Tasks {
		if defaultTask && task.Default {
			return ErrDuplicateDefault
		}

		if task.Default {
			defaultTask = task.Default
		}

		if len(task.Deps) == 0 {
			continue
		}

		for _, dep := range task.Deps {
			t, ok := settings.Tasks[strings.TrimSpace(dep)]
			if !ok {
				return errors.New(fmt.Sprintf("dep %s in %s does not exists", dep, name))
			}
			task.AddDep(t)
		}
	}

	s.Version = settings.Version
	s.Author = settings.Author
	s.Project = settings.Project
	s.Tasks = settings.Tasks
	s.Vars = settings.Vars

	return nil
}
