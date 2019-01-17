package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
}

var Test = &cobra.Command{
	Use:   "test",
	Short: "tests a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating project")

		return
	},
}
