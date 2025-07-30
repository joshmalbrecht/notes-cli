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
	// Disable the default 'completion' command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the notes-cli version",
	Long:  `Print the notes-cli version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Version)
	},
}
