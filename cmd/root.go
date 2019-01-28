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

	// init config
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
	addTaskCommands(root)

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func addTaskCommands(root *cobra.Command) error {
	// configure path
	cwd, err := cfg.Cwd()
	if err != nil {
		return err
	}

	// add create task
	root.AddCommand(generateCreate(cwd))

	// settings opts
	sopts := []s.Opt{func(o *s.Opts) {
		o.File = path.Join(cwd, cfg.File)
		o.FileMode = cfg.FileMode
	}}

	// new settings
	settings := config.NewSettings()
	ss := s.New(sopts...)
	if err := ss.Read(&settings); err != nil && err != os.ErrNotExist {
		return err
	}

	// attach tasks
	for use, task := range settings.Tasks {
		root.AddCommand(generateTask(use, task))
	}

	return nil
}

func initConfig() {}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
