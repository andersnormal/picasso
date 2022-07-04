package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/andersnormal/picasso/pkg/config"
	"github.com/andersnormal/picasso/pkg/executr"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run task from your task file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// create root context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		s, err := cfg.LoadSpec()
		if err != nil {
			return err
		}

		tt := s.Default()
		if len(args) == 0 && len(tt) == 0 {
			return config.ErrNoDefaultTask
		}

		tt, err = s.Find(args)
		if err != nil {
			return err
		}

		env := executr.Env{}
		for _, v := range cfg.RunConfig.Env {
			k := strings.Split(v, "=")
			if len(k) != 2 {
				return fmt.Errorf("invalid env var: %s", v)
			}
			env[k[0]] = k[1]
		}

		exec := executr.New(
			executr.WithEnv(env),
			executr.WithTimeout(cfg.RunConfig.Timeout),
		)
		for _, t := range tt {
			if err := exec.Run(ctx, t); err != nil {
				return err
			}
		}

		return nil
	},
}
