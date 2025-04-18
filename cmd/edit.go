package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joshmalbrecht/note/internal/config"
	"github.com/joshmalbrecht/note/internal/notes"
	"github.com/joshmalbrecht/note/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing note",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configurations, err := config.Get()
		if err != nil {
			fmt.Printf("error reading configurations: %v", err)
			os.Exit(1)
		}

		existingNotes, err := notes.List(configurations.NotesLocation)
		if err != nil {
			fmt.Printf("error to retrieving notes: %v", err)
			os.Exit(1)
		}

		model := ui.InitialModel(existingNotes)
		program := tea.NewProgram(model)
		if _, err := program.Run(); err != nil {
			fmt.Printf("error running selection ui: %v", err)
			os.Exit(1)
		}

		// Find the selected filename
		var index int
		for i, val := range model.Selected {
			if val == struct{}{} {
				index = i
				break
			}
		}

		var filename = model.Choices[index]
		err = notes.Edit(filename)
		if err != nil {
			fmt.Printf("error editing note: %v", err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
