package main

import (
	"fmt"

	"github.com/andersnormal/picasso/pkg/proto"
)

func main() {
	proto.Options{}.Run(func(p *proto.Plugin) error {
		// dummy log spec
		return fmt.Errorf("hello he %s", "help")
	})
}
