package settings

import (
	"os"
)

type Settings interface {
	Read(in interface{}) error
	Write(out interface{}) error
}

type settings struct {
	opts *Opts
}

type Opt func(*Opts)

type Opts struct {
	File     string
	FileMode os.FileMode
}
