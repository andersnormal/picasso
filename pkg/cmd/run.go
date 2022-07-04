package cmd

import (
	"fmt"

	"github.com/andersnormal/picasso/pkg/config"
	"github.com/spf13/cobra"
)

var run = &cobra.Command{
	Use:   "run [command]",
	Short: "Run task from your task file",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		fmt.Println(tt)

		return nil
	},
}
