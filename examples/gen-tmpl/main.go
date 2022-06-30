package main

import (
	"fmt"

	"github.com/andersnormal/picasso/pkg/plugin"
)

func main() {
	plugin.Options{}.Run(func(p *plugin.Plugin) error {
		// dummy log spec
		return fmt.Errorf("hello he %s", "help")
	})
}
