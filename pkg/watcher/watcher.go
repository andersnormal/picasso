package watcher

import (
	"context"
	"path"
	"time"

	"github.com/andersnormal/picasso/pkg/spec"

	"github.com/fsnotify/fsnotify"
)

// Watcher ...
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
	task spec.Task

	fs *fsnotify.Watcher
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Paths []string
	Cwd   string
}

// New ...
func New(task spec.Task, opts ...Opt) Watcher {
	options := &Opts{}
	options.Configure(opts...)

	var w = new(watcher)
	w.opts = options
	w.task = task

	return w
}

// Errors ...
func (w *watcher) Errors() <-chan error {
	return w.fs.Errors
}

// Events ...
func (w *watcher) Events() <-chan fsnotify.Event {
	return w.fs.Events
}

// Reload ...
func (w *watcher) Reload(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)

	fs, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	w.fs = fs

	for _, p := range w.opts.Paths {
		if err := fs.Add(path.Join(w.opts.Cwd, p)); err != nil {
			return err
		}
	}

	for range ticker.C {
		return nil
	}

	return nil
}

func (w *watcher) Close() {
	if w.fs == nil {
		return
	}

	w.fs.Close()
}

func (w *watcher) Stop() {
	w.stop <- true
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}
