package plugin

import (
	p "github.com/hashicorp/go-plugin"

	"github.com/andersnormal/picasso/pkg/proto"
)

const (
	PluginName             = "provider"
	DefaultProtocolVersion = 1
)

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
func Serve() {
	p.Serve(&p.ServeConfig{
		GRPCServer: p.DefaultGRPCServer,
	})
}
