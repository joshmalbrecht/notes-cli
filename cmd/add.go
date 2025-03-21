package cmd

import (
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
			println("Must provide only one argument")
			return
		}

		configurations, err := config.Get()
		if err != nil {
			println("unable to read config: " + err.Error())
		}

		title := args[0]

		created, err := notes.Create(configurations.NotesLocation, title)
		if err != nil {
			println("Unable to create note: " + err.Error())
			return
		}

		println(created + " has been created")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
