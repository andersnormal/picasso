package cmd

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

func init() {
}

var Create = &cobra.Command{
	Use:   "create",
	Short: "creates a new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating project")

		box := packr.New("myBox", "../templates")

		s, err := box.FindString("_README.md")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)

		return
	},
}
