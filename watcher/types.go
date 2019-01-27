package watcher

import (
	"context"

	"github.com/andersnormal/picasso/task"

	"github.com/fsnotify/fsnotify"
)

type Watcher interface {
	Reload(ctx context.Context) error
	Errors() <-chan error
	Events() <-chan fsnotify.Event
	Close()
	Stop()
}

type watcher struct {
	opts *Opts
	stop chan bool
	task *task.Task

	fs *fsnotify.Watcher
}

type Opt func(*Opts)

type Opts struct {
	Paths []string
	Cwd   string
}
