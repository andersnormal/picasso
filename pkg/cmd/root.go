package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/andersnormal/picasso/pkg/config"
	s "github.com/andersnormal/picasso/pkg/settings"
	"github.com/andersnormal/picasso/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg *config.Config

var root = &cobra.Command{
	Use:     "picasso",
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.ErrNoDefaultTask
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
	addFlags()

	// set default formatter
	log.SetFormatter(&log.TextFormatter{})

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// add run commands
	if err := addTaskCommands(run); err != nil {
		log.Fatal(err)
	}

	// add sub-commands
	root.AddCommand(run)
	root.AddCommand(initCmd)

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func addTaskCommands(root *cobra.Command) error {
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
	if err := ss.Read(&settings); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// attach tasks
	for use, task := range settings.Tasks {
		if task.Disable {
			continue
		}

		t := generateTask(use, task)
		if task.Default {
			root.RunE = t.RunE
		}
		root.AddCommand(generateTask(use, task))
	}

	return nil
}

func initConfig() {
	// unmarshal to config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(fmt.Sprintf("cannot unmarshal config: %v", err))
	}

	// setup logger
	cfg.SetupLogger()
}

func addFlags() {
	initCmd.Flags().StringVarP(&cfg.InitConfig.Folder, "folder", "f", cfg.InitConfig.Folder, "folder")
	initCmd.Flags().StringVarP(&cfg.InitConfig.URL, "url", "u", cfg.InitConfig.Folder, "url of archive")
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Error(err)

		if err != config.ErrNoDefaultTask {
			os.Exit(0)
		}

		os.Exit(1)
	}
}
