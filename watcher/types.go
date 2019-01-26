package watcher

import (
	"github.com/fsnotify/fsnotify"
)

type Watcher interface {
	Watch() error
	Close()
	Stop()
}

type watcher struct {
	opts *Opts
	stop chan bool

	fs *fsnotify.Watcher
}

type Opt func(*Opts)

type Opts struct {
	Paths []string
	Cwd   string
}
