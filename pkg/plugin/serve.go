package plugin

import (
	p "github.com/hashicorp/go-plugin"

	"github.com/andersnormal/picasso/pkg/proto"
)

const (
	PluginName             = "plugin"
	DefaultProtocolVersion = 1
)

var VersionedPlugins = map[int]p.PluginSet{
	1: {
		"plugin": &GRPCTaskPlugin{},
	},
}

// Handshake ...
var Handshake = p.HandshakeConfig{
	ProtocolVersion: DefaultProtocolVersion,

	MagicCookieKey:   "PICASSO_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "iaafij5485d5utqh",
}

// GRPCPluginFunc ...
type GRPCPluginFunc func() proto.PluginServer

// ServeOpts ...
type ServeOpts struct {
	GRPCPluginFunc GRPCPluginFunc
}

// Serve ...
func Serve(opts *ServeOpts) {
	p.Serve(&p.ServeConfig{
		GRPCServer:       p.DefaultGRPCServer,
		HandshakeConfig:  Handshake,
		VersionedPlugins: pluginSet(opts),
	})
}

func pluginSet(opts *ServeOpts) map[int]p.PluginSet {
	plugins := map[int]p.PluginSet{}

	// add the new protocol versions if they're configured
	if opts.GRPCPluginFunc != nil {
		plugins[1] = p.PluginSet{
			"plugin": &GRPCTaskPlugin{
				GRPCPlugin: opts.GRPCPluginFunc,
			},
		}
	}

	return plugins
}
