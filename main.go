package main

import (
	"math/rand"
	"time"

	"github.com/andersnormal/picasso/pkg/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
