package main

//go:generate echo "Generating..."
//go:generate protoc -I. --go_out=paths=source_relative:. pkg/plugin/plugin.proto
