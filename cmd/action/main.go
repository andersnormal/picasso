package main

import (
	"github.com/andersnormal/picasso/pkg/config"

	githubactions "github.com/sethvargo/go-githubactions"
)

var action *githubactions.Action

func run() error {
	action = githubactions.New()

	c := config.New()
	err := c.InitActionConfig(action)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		action.Fatalf("%v", err)
	}
}
