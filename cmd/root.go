package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andersnormal/picasso/config"

	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var (
	settings = config.NewSettings()
)

var rootCmd = &cobra.Command{
	Use:     "picasso",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func init() {
	// seed
	rand.Seed(time.Now().UnixNano())

	// create config
	cfg = config.New()

	// setup logger
	cfg.SetupLogger()

	// add flags
	cfg.AddFlags(rootCmd)

	// silence on the root cmd
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	// add commands
	rootCmd.AddCommand(Create)
	rootCmd.AddCommand(Build)
	rootCmd.AddCommand(Test)
	rootCmd.AddCommand(Watch)

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
