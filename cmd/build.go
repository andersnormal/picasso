package cmd

import (
	"github.com/spf13/cobra"
)

func init() {}

var Build = &cobra.Command{
	Use:   "build",
	Short: "build a new project",
	Run: func(cmd *cobra.Command, args []string) {

		return
	},
}
