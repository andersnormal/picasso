package main

import (
	"fmt"

	"github.com/andersnormal/picasso/pkg/plugin"
)

func main() {
	plugin.Options{}.Run(func(gen *plugin.Plugin) error {
		// dummy log spec
		fmt.Println(gen.Spec)

		return nil
	})
}
