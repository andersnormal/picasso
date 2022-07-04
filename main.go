package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/andersnormal/picasso/pkg/config"
	"github.com/andersnormal/picasso/pkg/executr"

	"github.com/spf13/pflag"
)

const usage = `Usage: picasso [-cfvsd] [--config] [--force] [--verbose] [--silent] [--dry] [task...]

'''
spec: 	 1
tasks:
  test:
    cmd:
      - go test -v ./...
'''

Options:
`

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	cfg := config.New()

	err := cfg.InitDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	pflag.Usage = func() {
		log.Print(usage)
		pflag.PrintDefaults()
	}

	pflag.BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	pflag.BoolVarP(&cfg.Flags.Help, "help", "h", cfg.Flags.Help, "show help")
	pflag.BoolVarP(&cfg.Flags.Init, "init", "i", cfg.Flags.Init, "init config")
	pflag.BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
	pflag.BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	pflag.BoolVarP(&cfg.Flags.Silent, "silent", "s", cfg.Flags.Silent, "silent mode")
	pflag.StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	pflag.StringSliceVarP(&cfg.Flags.Env, "env", "e", cfg.Flags.Env, "environment variables")
	pflag.Parse()

	if cfg.Flags.Help {
		pflag.Usage()
		os.Exit(0)
	}

	s, err := cfg.LoadSpec()
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range s.Tasks {
		log.Printf("%s", t.Name)
	}

	pflag.Parse()

	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	tt, err := s.Find(args)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	exec := executr.New(
		executr.WithTimeout(cfg.RunConfig.Timeout),
	)

	for _, t := range tt {
		if err := exec.Run(ctx, t); err != nil {
			log.Fatal(err)
		}
	}
}

func parseArgs() ([]string, error) {
	args := pflag.Args()
	dash := pflag.CommandLine.ArgsLenAtDash()

	if dash == -1 {
		return args, nil
	}

	return args[:dash], nil
}
