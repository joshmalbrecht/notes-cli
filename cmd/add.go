package cmd

import (
	"fmt"
	"os"

	"github.com/joshmalbrecht/note/internal/config"
	"github.com/joshmalbrecht/note/internal/notes"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new note",
	Long:  `Opens the vi text editor to create a new note. After adding your text and closing the editor, note will be saved in the configured directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Must provide only one argument")
			os.Exit(1)
		}

		configurations, err := config.Get()
		if err != nil {
			fmt.Println("unable to read config: " + err.Error())
			os.Exit(1)
		}

		title := args[0]

		filename, created, err := notes.Create(configurations.NotesLocation, title, configurations.FileExtension)
		if err != nil {
			fmt.Println("Unable to create note: " + err.Error())
			os.Exit(1)
		}

		if created {
			fmt.Println(filename + " has been created")
		} else {
			fmt.Println("Empty note was not created")
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
