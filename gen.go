package main

//go:generate echo "Generating..."
//go:generate protoc -I. --go_out=paths=source_relative:. pkg/proto/plugin.proto
//go:generate protoc -I. --go_out=paths=source_relative:. pkg/proto/spec.proto
//go:generate protoc -I. --go_out=paths=source_relative:. pkg/proto/generator.proto
