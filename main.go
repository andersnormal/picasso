package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/debug"
	"time"

	"github.com/andersnormal/picasso/pkg/config"
	"github.com/andersnormal/picasso/pkg/executr"
	"github.com/andersnormal/picasso/pkg/spec"
	"github.com/andersnormal/pkg/utils"

	"gopkg.in/yaml.v3"

	"github.com/spf13/pflag"
)

var (
	version = ""
)

const usage = `Usage: picasso [-cfglvsd] [--config] [--force] [--generator] [--list] [--verbose] [--silent] [--dry] [--validate] [--init] [--version] [task...] 

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
	pflag.BoolVar(&cfg.Flags.Init, "init", cfg.Flags.Init, "init config")
	pflag.BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
	pflag.BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	pflag.BoolVarP(&cfg.Flags.Silent, "silent", "s", cfg.Flags.Silent, "silent mode")
	pflag.StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	pflag.StringSliceVarP(&cfg.Flags.Env, "env", "e", cfg.Flags.Env, "environment variables")
	pflag.StringVarP(&cfg.Flags.Generator, "generator", "g", cfg.Flags.Generator, "generator")
	pflag.BoolVarP(&cfg.Flags.Validate, "validate", "V", cfg.Flags.Validate, "validate config")
	pflag.BoolVarP(&cfg.Flags.List, "list", "l", cfg.Flags.List, "list tasks")
	pflag.DurationVarP(&cfg.Flags.Timeout, "timeout", "t", time.Second*300, "timeout")
	pflag.BoolVar(&cfg.Flags.Version, "version", cfg.Flags.Version, "version")
	pflag.Parse()

	if cfg.Flags.Version {
		fmt.Printf("%s\n", getVersion())
		return
	}

	if cfg.Flags.Help {
		pflag.Usage()
		os.Exit(0)
	}

	s, err := cfg.LoadSpec()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Flags.Validate {
		err = s.Validate()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if cfg.Flags.List {
		for k, t := range s.Tasks {
			log.Printf("%s (%s)", k, t.Description)
		}
		os.Exit(0)
	}

	if cfg.Flags.Init {
		s := &spec.Spec{
			Spec:    1,
			Version: "0.0.1",
			Tasks:   map[string]spec.Task{},
		}

		b, err := yaml.Marshal(&s)
		if err != nil {
			log.Fatal(err)
		}

		ok, err := utils.FileExists(cfg.File)
		if err != nil {
			log.Fatal(err)
		}

		if ok && !cfg.Flags.Force {
			log.Fatalf("%s already exists, use --force to overwrite", cfg.File)
		}

		f, err := os.Create(cfg.File)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.Write(b)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	tasks := s.Default()

	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 0 && len(tasks) == 0 {
		log.Fatal("no default task")
	}

	tt, err := s.Find(args)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	exec := executr.New(
		executr.WithTimeout(cfg.Flags.Timeout),
		executr.WithStderr(cfg.Stderr),
		executr.WithStdin(cfg.Stdin),
		executr.WithStdout(cfg.Stdout),
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

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version = info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}
