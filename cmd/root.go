package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andersnormal/picasso/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg *config.Config
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

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// reading config file
	if cfg.CfgFile != "" {
		viper.SetConfigFile(cfg.CfgFile)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(dir)
		viper.SetConfigName(config.ConfigFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config: %v", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
