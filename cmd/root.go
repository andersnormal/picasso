package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/andersnormal/picasso/config"
	s "github.com/andersnormal/picasso/settings"

	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var (
	settings = config.NewSettings()
)

var root = &cobra.Command{
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
	cfg.AddFlags(root)

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// add commands
	addComamnds(root)

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func addComamnds(root *cobra.Command) error {
	// configure path
	cwd, err := cfg.Cwd()
	if err != nil {
		return err
	}

	// settings opts
	sopts := []s.Opt{func(o *s.Opts) {
		o.File = path.Join(cwd, cfg.File)
		o.FileMode = cfg.FileMode
	}}

	// new settings
	settings := config.NewSettings()
	ss := s.New(sopts...)
	if err := ss.Read(&settings); err != nil {
		return err
	}

	for use, task := range settings.Tasks {
		root.AddCommand(generateTask(use, task))
	}

	// add create not programmatically
	root.AddCommand(Create)

	return nil
}

func initConfig() {
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
