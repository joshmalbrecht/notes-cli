package cmd

import (
	"fmt"

	"github.com/joshmalbrecht/note/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the notes configuration directory and file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Initialize()
		if err != nil {
			fmt.Println("unable to initialize: " + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
