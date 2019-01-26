package watcher

import (
	"fmt"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
)

func New(opts ...Opt) Watcher {
	options := &Opts{}

	var w = new(watcher)
	w.opts = options

	configure(w, opts...)

	return w
}

func (w *watcher) Watch() error {
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

	for {
		select {
		case <-ticker.C:
		case events := <-w.fs.Events:
			fmt.Println(events)
		default:
		}
	}
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

func configure(w *watcher, opts ...Opt) error {
	for _, o := range opts {
		o(w.opts)
	}

	w.stop = make(chan bool)

	return nil
}
