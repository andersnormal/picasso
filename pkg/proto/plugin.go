package proto

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

// Run ...
func (o Options) Run(f func(*Plugin) error) {
	if err := run(o, f); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

func run(opts Options, f func(*Plugin) error) error {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	req := &PluginRequest{}
	err = proto.Unmarshal(in, req)
	if err != nil {
		return err
	}

	p, err := opts.New(req)
	if err != nil {
		return err
	}

	if err := f(p); err != nil {
		p.Error(err)
	}

	res := p.Response()
	out, err := proto.Marshal(res)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// Options ...
type Options struct {
	// ParamFunc ...
	paramFunc func(name, value string) error
}

// Plugin ...
type Plugin struct {
	Parameters map[string]string
	Version    string
	Spec       string

	opts Options
	err  error
}

// New returns a new Plugin.
func (opts Options) New(req *PluginRequest) (*Plugin, error) {
	gen := &Plugin{
		Parameters: req.GetParameters(),
		Version:    req.GetVersion(),
		Spec:       req.GetSpec(),
	}

	return gen, nil
}

// Error ...
func (p *Plugin) Error(err error) {
	if p.err == nil {
		p.err = err
	}
}

// Response ...
func (p *Plugin) Response() *PluginResponse {
	res := &PluginResponse{}

	if p.err != nil {
		res.Error = p.err.Error()
	}

	return res
}
