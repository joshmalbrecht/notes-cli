package cmd

import (
	"fmt"
	"os"

	"github.com/joshmalbrecht/note/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "A note taking CLI",
	Long:  `A CLI that quickly creates and organizes notes`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the notes-cli version",
	Long:  `Print the notes-cli version`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := config.GetVersion()
		if err != nil {
			fmt.Println("unable to retrieve version: " + err.Error())
		}

		fmt.Println(version)
	},
}
