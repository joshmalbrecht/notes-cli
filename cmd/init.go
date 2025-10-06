package cmd

import (
	"fmt"
	"os"

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
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
