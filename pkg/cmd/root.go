package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andersnormal/picasso/pkg/config"
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

	err := cfg.InitDefaultConfig()
	if err != nil {
		panic(err)
	}

	// add flags
	cfg.AddFlags(root)
	addFlags()

	// set default formatter
	log.SetFormatter(&log.TextFormatter{})

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// add sub-commands
	root.AddCommand(runCmd)
	root.AddCommand(initCmd)
	root.AddCommand(validateCmd)

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// unmarshal to config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(fmt.Sprintf("cannot unmarshal config: %v", err))
	}

	// setup logger
	_ = cfg.SetupLogger()
}

func addFlags() {
	initCmd.Flags().BoolVar(&cfg.InitConfig.ArchiveMode, "archive", cfg.InitConfig.ArchiveMode, "url is an archive")
	initCmd.Flags().StringVarP(&cfg.InitConfig.Folder, "folder", "f", cfg.InitConfig.Folder, "folder")
	initCmd.Flags().StringVarP(&cfg.InitConfig.URL, "url", "u", cfg.InitConfig.Folder, "url of archive")

	runCmd.Flags().StringSliceVarP(&cfg.RunConfig.Env, "env", "e", cfg.RunConfig.Env, "environment variables")
	runCmd.Flags().DurationVar(&cfg.RunConfig.Timeout, "timeout", cfg.RunConfig.Timeout, "timeout")
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
