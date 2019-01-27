package config

import (
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
	var settings struct {
		Version string
		Author  string
		Project string
		Tasks   task.Tasks
		Vars    Vars
	}
	if err := unmarshal(&settings); err != nil {
		return err
	}

	s.Version = settings.Version
	s.Author = settings.Author
	s.Project = settings.Project
	s.Tasks = settings.Tasks
	s.Vars = settings.Vars

	return nil
}
