package main

import (
	"math/rand"
	"time"

	"github.com/andersnormal/picasso/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
