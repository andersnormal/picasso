package main

import (
	"context"
	"fmt"

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
func (s *server) Execute(ctx context.Context, req *proto.Execute_Request) (*proto.Execute_Response, error) {
	fmt.Println("Execute")

	return &proto.Execute_Response{}, nil
}

// Stop ...
func (s *server) Stop(ctx context.Context, req *proto.Stop_Request) (*proto.Stop_Response, error) {
	fmt.Println("Stop")

	return &proto.Stop_Response{}, nil
}
