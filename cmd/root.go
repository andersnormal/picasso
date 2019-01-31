package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/andersnormal/picasso/config"
	s "github.com/andersnormal/picasso/settings"
	"github.com/andersnormal/picasso/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var root = &cobra.Command{
	Use:     "picasso",
	Version: version.Version,
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

	// add flags
	cfg.AddFlags(root)

	// set default formatter
	log.SetFormatter(&log.TextFormatter{})

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// add commands
	if err := addTaskCommands(root); err != nil {
		log.Fatal(err)
	}

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
	if err := ss.Read(&settings); err != nil {
		return err
	}

	// attach tasks
	for use, task := range settings.Tasks {
		root.AddCommand(generateTask(use, task))
	}

	return nil
}

func initConfig() {
	// setup logger
	cfg.SetupLogger()
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
