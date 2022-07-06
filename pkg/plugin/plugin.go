package plugin

import (
	"context"
	"errors"
	"os/exec"

	"github.com/hashicorp/go-plugin"

	"github.com/andersnormal/picasso/pkg/proto"
)

// Meta ...
type Meta struct {
	// Path ...
	Path string
}

func (m *Meta) factory() Factory {
	return pluginFactory(m)
}

// GRPCPlugin ...
type GRPCPlugin struct {
	PluginClient *plugin.Client

	ctx    context.Context
	client proto.PluginClient
}

// Stop ...
func (p *GRPCPlugin) Stop() error {
	resp, err := p.client.Stop(p.ctx, new(proto.Stop_Request))
	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// Start ...
func (p *GRPCPlugin) Close() error {
	if p.PluginClient != nil {
		return nil
	}

	p.PluginClient.Kill()
	return nil
}

// Factory ...
type Factory func() (Plugin, error)

// Plugin ...
type Plugin interface {
	// Stop ...
	Stop() error
	// Close ...
	Close() error
}

func pluginFactory(meta *Meta) Factory {
	return func() (Plugin, error) {
		cfg := &plugin.ClientConfig{
			Cmd: exec.Command(meta.Path),
		}
		client := plugin.NewClient(cfg)

		rpc, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpc.Dispense("help")
		if err != nil {
			return nil, err
		}

		p := raw.(*GRPCPlugin)
		p.PluginClient = client

		return p, nil
	}
}
