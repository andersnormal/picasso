package main

import (
	"context"

	"github.com/andersnormal/picasso/pkg/plugin"
	"github.com/andersnormal/picasso/pkg/proto"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCPluginFunc: func() proto.PluginServer {
			return &server{}
		},
	})
}

type server struct {
	proto.UnimplementedPluginServer
}

// Start ...
func (s *server) Start(ctx context.Context, req *proto.Start_Request) (*proto.Start_Response, error) {
	return nil, nil
}

// Stop ...
func (s *server) Stop(ctx context.Context, req *proto.Stop_Request) (*proto.Stop_Response, error) {
	return nil, nil
}
