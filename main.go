package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/andersnormal/picasso/pkg/config"
	"github.com/andersnormal/picasso/pkg/executr"
	"github.com/andersnormal/picasso/pkg/plugin"
	"github.com/andersnormal/picasso/pkg/spec"
	"github.com/andersnormal/pkg/utils"

	"github.com/spf13/pflag"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

var (
	version = ""
)

const usage = `Usage: picasso [-cfglvsdpw] [--config] [--force] [--generator] [--list] [--verbose] [--silent] [--dry] [--plugin] [--watch] [--validate] [--var] [--init] [--version] [task...] 

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
	pflag.StringVarP(&cfg.Flags.Plugin, "plugin", "p", cfg.Flags.Plugin, "plugin")
	pflag.BoolVarP(&cfg.Flags.Validate, "validate", "V", cfg.Flags.Validate, "validate config")
	pflag.BoolVarP(&cfg.Flags.List, "list", "l", cfg.Flags.List, "list tasks")
	pflag.DurationVarP(&cfg.Flags.Timeout, "timeout", "t", time.Second*300, "timeout")
	pflag.BoolVar(&cfg.Flags.Version, "version", cfg.Flags.Version, "version")
	pflag.StringSliceVar(&cfg.Flags.Vars, "var", cfg.Flags.Vars, "variables")
	pflag.BoolVarP(&cfg.Flags.Watch, "watch", "w", cfg.Flags.Watch, "watch")
	pflag.Parse()

	if cfg.Flags.Verbose {
		start := time.Now()
		defer func() { log.Printf("time: %s", time.Since(start)) }()
	}

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := map[string]string{}
	for _, v := range cfg.Flags.Vars {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			params[kv[0]] = ""
			continue
		}
		params[kv[0]] = kv[1]
	}

	if cfg.Flags.Plugin != "" {
		m := &plugin.Meta{Path: cfg.Flags.Plugin}
		f := m.Factory()

		p, err := f()
		if err != nil {
			log.Fatal(err)
		}
		defer p.Close()

		pp := s.Vars
		maps.Copy(pp, params)

		resp, err := p.Execute(plugin.ExecuteRequest{
			Parameters: pp,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)

		err = p.Stop()
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

	exec := executr.New(
		executr.WithTimeout(cfg.Flags.Timeout),
		executr.WithStderr(cfg.Stderr),
		executr.WithStdin(cfg.Stdin),
		executr.WithStdout(cfg.Stdout),
	)

	for _, task := range tt {
		pp := s.Vars
		maps.Copy(pp, task.Vars)
		maps.Copy(pp, params)
		task.Vars = pp

		if task.Disabled {
			continue
		}

		if err := exec.Run(ctx, task, cfg.Flags.Watch); err != nil {
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
