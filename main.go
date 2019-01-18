package main

import (
	"math/rand"
	"time"

	"github.com/andersnormal/picasso/cmd"
	"github.com/andersnormal/picasso/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "picasso",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())

	// silence on the root cmd
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	// setup logger
	config.C.SetupLogger()

	// add flags
	config.C.AddFlags(rootCmd)

	// add commands
	rootCmd.AddCommand(cmd.Create)
	rootCmd.AddCommand(cmd.Build)
	rootCmd.AddCommand(cmd.Test)
}

func main() {
	rootCmd.Execute()
}
